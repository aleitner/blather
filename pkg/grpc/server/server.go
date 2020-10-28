package server

import (
	"google.golang.org/grpc/metadata"
	"io"
	"strings"
	"sync"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
	log "github.com/sirupsen/logrus"
)

type CallServer struct {
	logger *log.Logger
}

func NewCallServer(logger *log.Logger) call.PhoneServer {
	return &CallServer{
		logger: logger,
	}
}

func (cs *CallServer) Call(stream call.Phone_CallServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		cs.logger.Error("Failed to retrieve incoming context")
	}

	clientID := md.Get("client-id")

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
			cs.logger.Infof("%s: %s", clientID, data.GetAudioData()[:data.GetLength()])
		}

		wg.Done()
	}()
	wg.Add(1)

	// Send data
	go func() {
		// TODO: mic input
		r := strings.NewReader("Howdy from the server")

		for {
			//TODO: Get mic audio. encompass the following lines into a Mic.Read()
			buf := make([]byte, 4)
			_, err := r.Read(buf)
			if err != nil {
				if err != io.EOF {
					cs.logger.Error(err.Error())
				}

				break
			}

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
