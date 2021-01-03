package server

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/aleitner/blather/internal/utils"
	"github.com/aleitner/blather/pkg/userid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	_ "google.golang.org/grpc/encoding/gzip"

	"github.com/aleitner/blather/pkg/forwarder"
	"github.com/aleitner/blather/pkg/protobuf"
	log "github.com/sirupsen/logrus"
)

// BlatherServer forwards call data to all the clients
type BlatherServer struct {
	logger *log.Logger
	rooms  map[string]*forwarder.Forwarder
}

// NewBlatherServer
func NewBlatherServer(logger *log.Logger) blatherpb.PhoneServer {
	return &BlatherServer{
		logger: logger,
		rooms:  make(map[string]*forwarder.Forwarder),
	}
}

func RegisterBlatherServer(registrar grpc.ServiceRegistrar, server blatherpb.PhoneServer) {
	blatherpb.RegisterPhoneServer(registrar, server)
}

// Call gets a stream of audio data from the client and forwards it to the other clients
func (bs *BlatherServer) Call(stream blatherpb.Phone_CallServer) error {
	// Get id from client
	md, err := expandMetaData(stream.Context())
	if err != nil {
		return fmt.Errorf("Failed to retrieve incoming metadata: %s", err.Error())
	}

	room := bs.rooms[md.RoomID]
	if room == nil {
		return fmt.Errorf("Room %s is not a valid room id", md.RoomID)
	}

	bs.logger.Infof("%s has joined room %s", md.ClientID.String(), md.RoomID)

	// Create forwarder for client
	room.Add(md.ClientID, stream)
	defer room.Delete(md.ClientID)

	var wg sync.WaitGroup // NB: we can probably just use a channel here

	// Receive data
	go func() {
		for {
			// Read Data
			data, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					bs.logger.Errorf("%s, %s, %s", md.ClientID.String(), md.RoomID, err.Error())
				}

				bs.logger.Infof("%s has left room %s", md.ClientID.String(), md.RoomID)
				break
			}

			// Forward the data to the other clients
			room.Forward(md.ClientID, data)
		}

		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	return nil
}

func (bs *BlatherServer) CreateRoom(ctx context.Context, req *blatherpb.CreateRoomReq) (*blatherpb.CreateRoomResp, error) {
	f := forwarder.NewForwarder()

	for i := 0; i < 1; i++ {
		rand.Seed(time.Now().UnixNano())
		roomID := utils.RandSeq(6)

		_, ok := bs.rooms[roomID]
		if !ok {
			bs.rooms[roomID] = f
			bs.logger.Infof("Created room %s", roomID)

			return &blatherpb.CreateRoomResp{
				Id: roomID,
			}, nil
		}

		if i == 10 {
			return nil, fmt.Errorf("Failed to create a valid room id")
		}
	}

	return nil, nil
}

func (bs *BlatherServer) UpdateSettings(ctx context.Context, userdata *blatherpb.UserSettingsData) (*blatherpb.UserSettingsResponse, error) {
	return nil, nil
}

type MD struct {
	ClientID userid.ID
	RoomID   string
}

func expandMetaData(ctx context.Context) (*MD, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to retrieve incoming context")
	}

	IDAsString := md.Get("client-id")

	if len(IDAsString) <= 0 {
		return nil, fmt.Errorf("Failed to retrieve incoming id")
	}

	contactID, err := userid.FromString(IDAsString[0])
	if err != nil {
		return nil, fmt.Errorf("Failed to parse incoming id: %s", err.Error())
	}

	RoomID := md.Get("room-id")

	if len(RoomID) <= 0 {
		return nil, fmt.Errorf("Failed to retrieve incoming room id")
	}

	return &MD{
		ClientID: contactID,
		RoomID:   RoomID[0],
	}, nil
}
