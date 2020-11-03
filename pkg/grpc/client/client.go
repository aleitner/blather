package client

import (
	"context"
	"io"
	"strconv"
	"sync"

	"github.com/aleitner/spacialPhone/internal/muxer"
	call "github.com/aleitner/spacialPhone/internal/protobuf"
	"github.com/aleitner/spacialPhone/internal/utils"
	"github.com/aleitner/spacialPhone/pkg/user/coordinates"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CallClient interface {
	Call(ctx context.Context, audioInput beep.Streamer, format beep.Format) error
	CloseConn() error
}

type Client struct {
	id         int
	logger     *log.Logger
	route      call.PhoneClient
	conn       *grpc.ClientConn
	coordinate *coordinates.Coordinate
	muxer      *muxer.Muxer
}

func NewClient(id int, logger *log.Logger, conn *grpc.ClientConn) CallClient {
	return &Client{
		id:         id,
		logger:     logger,
		conn:       conn,
		route:      call.NewPhoneClient(conn),
		muxer:      muxer.NewMuxer(logger),
		coordinate: &coordinates.Coordinate{X: 0, Y: 0, Z: 0},
	}
}

func (client *Client) Call(ctx context.Context, audioInput beep.Streamer, format beep.Format) error {
	clientId := strconv.Itoa(client.id)
	md := metadata.Pairs("client-id", clientId)
	ctx = metadata.NewOutgoingContext(ctx, md)

	speaker.Init(format.SampleRate, 512)
	speaker.Play(client.muxer)

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
			numSamples, ok := audioInput.Stream(buf)
			if !ok {
				// server returns with nil
				if audioInput.Err() != nil {
					client.logger.Errorf("audio read fail: %s/n", err)
				}
				break
			}

			if err := stream.Send(&call.CallData{
				AudioData: &call.AudioData{
					AudioEncoding: "mp3",
					Samples:       utils.ToGRPCSampleRate(buf, numSamples),
					NumSamples:    int32(numSamples),
					Format: &call.Format{
						SampleRate:  uint32(sampleRate),
						NumChannels: uint32(format.NumChannels),
						Precision:   uint32(format.Precision),
					},
				},
				UserMetaData: &call.UserMetaData{
					Id:          uint64(client.id),
					Coordinates: client.coordinate.ToGRPC(),
				},
			}); err != nil {
				client.logger.Errorf("stream Send fail: %s/n", err)
			}
		}

		//if err := stream.CloseSend(); err != nil {
		//	client.logger.Errorf("close send fail: %s\n", err)
		//}

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
