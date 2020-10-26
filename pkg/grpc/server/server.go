package server

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
)

type CallServer struct {
	logger log.Logger
	id     int64
	// access to microphone
}

func NewCallServer() call.PhoneServer {
	return &CallServer{}
}

func (cs *CallServer) Call(stream call.Phone_CallServer) error {
	var wg sync.WaitGroup

	// Receive data
	go func() {
		for {
			// Read Data
			data, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					cs.logger.Printf("error: %s", err)
				}

				break
			}

			// TODO: Play parsed audio or mux with other audio channels
			fmt.Printf(string(data.GetAudioData()[:data.GetLength()]))
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
					cs.logger.Printf("error: %s", err)
				}

				break
			}

			data := &call.CallData{
				AudioEncoding: "bytes",
				AudioData:     buf,
				Length:        int64(len(buf)),
			}

			if err := stream.Send(data); err != nil {
				if err != io.EOF {
					cs.logger.Printf("error: %s", err)
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
