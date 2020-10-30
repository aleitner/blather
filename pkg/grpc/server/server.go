package server

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/aleitner/spacialPhone/internal/forwarder"

	"github.com/aleitner/spacialPhone/pkg/user/user_id"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
	log "github.com/sirupsen/logrus"
)

type CallServer struct {
	logger    *log.Logger
	forwarder *forwarder.Forwarder
}

func NewCallServer(logger *log.Logger) call.PhoneServer {
	return &CallServer{
		logger:    logger,
		forwarder: forwarder.NewForwarder(),
	}
}

func (cs *CallServer) Call(stream call.Phone_CallServer) error {
	// Get id from client
	clientID, err := user_id.NewIDFromMetaData(stream.Context())
	if err != nil {
		return fmt.Errorf("Failed to retrieve incoming client id")
	}

	// Create forwarder for client
	cs.forwarder.Add(clientID, stream)
	defer cs.forwarder.Delete(clientID)

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

			// Forward the data to the other clients
			cs.forwarder.ForwardTo(clientID, data)
		}

		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	return nil
}

func (cs *CallServer) UpdateSettings(context.Context, *call.UserSettingsData) (*call.UserSettingsResponse, error) {
	return nil, nil
}
