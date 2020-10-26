package server

import (
	"fmt"
	call "github.com/aleitner/spacialPhone/internal/protobuf"
	"io"
	"log"
	"sync"
)

type CallServer struct {
	logger log.Logger
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
			fmt.Println(data)
		}

		wg.Done()
	}()
	wg.Add(1)

	// Send data
	go func() {
		for {
			//TODO: Get mic audio. encompass the following lines into a Mic.Read()
			data := &call.CallData{
				AudioEncoding: "bytes",
				AudioData: []byte{1,2,3,4,5},
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
