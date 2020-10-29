package server

import (
	"fmt"
	"io"
	"sync"

	"github.com/aleitner/spacialPhone/internal/muxer"
	"google.golang.org/grpc/metadata"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
	log "github.com/sirupsen/logrus"
)

type CallServer struct {
	logger *log.Logger
	muxer  *muxer.Muxer
}

func NewCallServer(logger *log.Logger) call.PhoneServer {
	return &CallServer{
		logger: logger,
		muxer:  muxer.NewMuxer(),
	}
}

func (cs *CallServer) Call(stream call.Phone_CallServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("Failed to retrieve incoming context")
	}

	clientIDAsString := md.Get("client-id")
	if len(clientIDAsString) <= 0 {
		return fmt.Errorf("Failed to retrieve incoming client id")
	}

	clientID, err := muxer.NewID(clientIDAsString[0])
	if err != nil {
		return fmt.Errorf("Failed to retrieve incoming client id")
	}

	cs.muxer.Add(clientID)
	defer cs.muxer.Delete(clientID)

	var wg sync.WaitGroup

	// Receive data
	go func() {
		for {
			// Read Data
			data, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					cs.logger.Error(err.Error())
				}

				break
			}

			// TODO: Play parsed audio or mux with other audio channels
			cs.muxer.Produce(clientID, data.GetAudioData()[:data.GetLength()])
			cs.logger.Infof("%s: %s", clientID, data.GetAudioData()[:data.GetLength()])
		}

		wg.Done()
	}()
	wg.Add(1)

	// Send data
	go func() {
		for {
			//TODO: Get mic audio. encompass the following lines into a Mic.Read()
			buf := cs.muxer.Consume(clientID)

			data := &call.CallData{
				AudioEncoding: "bytes",
				AudioData:     buf,
				Length:        uint64(len(buf)),
			}

			if err := stream.Send(data); err != nil {
				if err != io.EOF {
					cs.logger.Error(err.Error())
				}

				break
			}
		}

		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	return nil
}
