package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"sync"

	"github.com/aleitner/blather/pkg/forwarder"
	"github.com/aleitner/blather/pkg/protobuf"
	"github.com/aleitner/blather/pkg/userid"
	log "github.com/sirupsen/logrus"
)

// BlatherServer forwards call data to all the clients
type BlatherServer struct {
	logger    *log.Logger
	forwarder *forwarder.Forwarder
}

// NewBlatherServer
func NewBlatherServer(logger *log.Logger) blatherpb.PhoneServer {
	return &BlatherServer{
		logger:    logger,
		forwarder: forwarder.NewForwarder(),
	}
}

func RegisterBlatherServer(registrar grpc.ServiceRegistrar, server blatherpb.PhoneServer) {
	blatherpb.RegisterPhoneServer(registrar, server)
}

// Call gets a stream of audio data from the client and forwards it to the other clients
func (bs *BlatherServer) Call(stream blatherpb.Phone_CallServer) error {
	// Get id from client
	clientID, err := userid.IDFromContextMetaData(stream.Context())
	if err != nil {
		return fmt.Errorf("Failed to retrieve incoming client id")
	}

	// Create forwarder for client
	bs.forwarder.Add(clientID, stream)
	defer bs.forwarder.Delete(clientID)

	var wg sync.WaitGroup // NB: we can probably just use a channel here

	// Receive data
	go func() {
		for {
			// Read Data
			data, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					bs.logger.Error(err.Error())
				}

				break
			}

			// Forward the data to the other clients
			bs.forwarder.Forward(clientID, data)
		}

		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	return nil
}

func (bs *BlatherServer) UpdateSettings(context.Context, *blatherpb.UserSettingsData) (*blatherpb.UserSettingsResponse, error) {
	return nil, nil
}
