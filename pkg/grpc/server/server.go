package server

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/aleitner/spacialPhone/internal/forwarder"
	call "github.com/aleitner/spacialPhone/internal/protobuf"
	"github.com/aleitner/spacialPhone/pkg/user/userid"
	log "github.com/sirupsen/logrus"
)

// CallServer forwards call data to all the clients
type CallServer struct {
	logger    *log.Logger
	forwarder *forwarder.Forwarder
}

// NewCallServer
func NewCallServer(logger *log.Logger) call.PhoneServer {
	return &CallServer{
		logger:    logger,
		forwarder: forwarder.NewForwarder(),
	}
}

// Call gets a stream of audio data from the client and forwards it to the other clients
func (cs *CallServer) Call(stream call.Phone_CallServer) error {
	// Get id from client
	clientID, err := userid.IDFromMetaData(stream.Context())
	if err != nil {
		return fmt.Errorf("Failed to retrieve incoming client id")
	}

	// Create forwarder for client
	cs.forwarder.Add(clientID, stream)
	defer cs.forwarder.Delete(clientID)

	var wg sync.WaitGroup // NB: we can probably just use a channel here

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
			cs.forwarder.Forward(clientID, data)
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
