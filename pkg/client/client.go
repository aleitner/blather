package client

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/aleitner/blather/internal/utils"
	"github.com/aleitner/blather/pkg/muxer"
	"github.com/aleitner/blather/pkg/protobuf"
	"github.com/faiface/beep"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CallClient interface {
	CreateRoom(ctx context.Context) (roomID string, err error)
	Call(ctx context.Context, room string, audioInput beep.Streamer, format beep.Format) error
	CloseConn() error
}

type Client struct {
	id         int
	logger     *log.Logger
	route      blatherpb.PhoneClient
	conn       *grpc.ClientConn
	muxer      *muxer.Muxer
	sr 		   beep.SampleRate
	quality    int
}

func NewClient(id int, logger *log.Logger, conn *grpc.ClientConn) CallClient {
	return &Client{
		id:         id,
		logger:     logger,
		conn:       conn,
		route:      blatherpb.NewPhoneClient(conn),
		muxer:      muxer.NewMuxer(),

		// Audio mixing
		sr: beep.SampleRate(44100),
		quality: 4,
	}
}

func (client *Client) CreateRoom(ctx context.Context) (roomID string, err error) {
	resp, err := client.route.CreateRoom(ctx, &blatherpb.CreateRoomReq{})
	if err != nil {
		return "", fmt.Errorf("Failed to create room: %s", err.Error())
	}

	return resp.GetId(), nil
}

func (client *Client) Call(ctx context.Context, room string, audioInput beep.Streamer, format beep.Format) error {
	clientId := strconv.Itoa(client.id)
	md := metadata.Pairs("client-id", clientId, "room-id", room)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Resample audio
	// NB: Perhaps we can determine the sample rate based on everyone's connections
	resampled := beep.Resample(client.quality, format.SampleRate, client.sr, audioInput)

	stream, err := client.route.Call(ctx)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	// Send data
	go func() {
		for {
			sampleRate := 512
			buf := make([][2]float64, sampleRate) // Optimal sending size is 16KiB-64KiB
			numSamples, ok := resampled.Stream(buf)
			if !ok {
				// server returns with nil
				if resampled.Err() != nil {
					client.logger.Errorf("audio read fail: %s/n", err)
				}
				break
			}

			if err := stream.Send(&blatherpb.CallData{
				AudioData: &blatherpb.AudioData{
					Samples:       utils.ToGRPCSampleRate(buf, numSamples),
					NumSamples:    uint32(numSamples),
				},
				UserId: uint64(client.id),
			}); err != nil {
				client.logger.Errorf("stream Send fail: %s/n", err)
			}
		}

		if err := stream.CloseSend(); err != nil {
			client.logger.Errorf("close send fail: %s\n", err)
		}

		wg.Done()
		client.logger.Errorf("Finished sending data...\n")
	}()
	wg.Add(1)

	// Receive data
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				client.logger.Errorf("stream Recv fail: %s/n", err)
			}

			// Add audio data to muxer
			client.muxer.Add(res)
		}

		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	return nil
}

// CloseConn closes conn
func (client *Client) CloseConn() error {
	return client.conn.Close()
}
