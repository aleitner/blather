package client

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/aleitner/blather/internal/utils"
	"github.com/aleitner/blather/pkg/coordinates"
	"github.com/aleitner/blather/pkg/muxer"
	"github.com/aleitner/blather/pkg/protobuf"
	"github.com/aleitner/blather/pkg/strmr"
	"github.com/aleitner/blather/pkg/userid"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
)

type CallClient interface {
	CreateRoom(ctx context.Context) (roomID string, err error)
	Call(ctx context.Context, room string, audioInput beep.Streamer, format beep.Format) error
	CloseConn() error
}

type Client struct {
	id           int
	logger       *log.Logger
	route        blatherpb.PhoneClient
	conn         *grpc.ClientConn
	Muxer        *muxer.Muxer
	resampleRate beep.SampleRate
	quality      int
	coordinates	 *coordinates.Coordinates
}

func NewClient(id int, logger *log.Logger, conn *grpc.ClientConn) *Client {
	return &Client{
		id:     id,
		logger: logger,
		conn:   conn,
		route:  blatherpb.NewPhoneClient(conn),
		Muxer:  muxer.NewMuxer(logger),

		// Audio mixing
		resampleRate: beep.SampleRate(11025),
		quality:      4,
		coordinates: &coordinates.Coordinates{},
	}
}

func (client *Client) CreateRoom(ctx context.Context) (roomID string, err error) {
	resp, err := client.route.CreateRoom(ctx, &blatherpb.CreateRoomReq{})
	if err != nil {
		return "", fmt.Errorf("Failed to create room: %s", err.Error())
	}

	return resp.GetId(), nil
}

func (client *Client) Call(ctx context.Context, room string, audioInput beep.Streamer, inputSampleRate int) error {
	clientId := strconv.Itoa(client.id)
	md := metadata.Pairs("client-id", clientId, "room-id", room)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Resample audio
	// NB: Perhaps we can determine the sample rate based on everyone's connections
	resampled := beep.Resample(client.quality, beep.SampleRate(inputSampleRate), client.resampleRate, audioInput)
	resampled.SetRatio(1)

	stream, err := client.route.Call(ctx, grpc.UseCompressor(gzip.Name))
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	// Send data
	go func() {
		for {
			buf := make([][2]float64, 2*1024) // Optimal sending size is 16KiB-64KiB

			numSamples, ok := resampled.Stream(buf)
			if !ok {
				// server returns with nil
				if resampled.Err() != nil {
					client.logger.Errorf("audio read fail: %s/n", err)
				}
				break
			}

			if numSamples == 0 {
				continue
			}

			if err := stream.Send(&blatherpb.CallData{
				UserId: uint64(client.id),
				AudioData: &blatherpb.AudioData{
					Samples:    utils.ToGRPCSampleRate(buf, numSamples),
					NumSamples: uint32(numSamples),
					SampleRate: uint32(inputSampleRate),
				},
				Coordinates: client.coordinates.ToGRPC(),
			}); err != nil {
				client.logger.Errorf("stream Send fail: %s/n", err)
				break
			}
		}

		if err := stream.CloseSend(); err != nil {
			client.logger.Errorf("close send fail: %s\n", err)
		}

		wg.Done()
		client.logger.Infof("Finished sending data...\n")
	}()
	wg.Add(1)

	// Receive data
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				client.logger.Errorf("stream Recv fail: %s/n", err)
				break
			}

			id, volumeStreamer := client.StreamerFromGRPC(res)
			if volumeStreamer == nil {
				client.logger.Errorf("stream Recv fail: %s/n", err)
				continue
			}

			// Add audio data to Muxer
			client.Muxer.Add(id, volumeStreamer)
		}

		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	return nil
}

func (client *Client) StreamerFromGRPC(data *blatherpb.CallData) (userid.ID, *effects.Volume) {
	audioData := data.GetAudioData()
	if audioData.GetNumSamples() == 0 {
		return 0, nil
	}

	grpcSamples := audioData.GetSamples()
	numSamples := int(audioData.GetNumSamples())
	id := userid.ID(data.GetUserId())

	samples := utils.ToSampleRate(grpcSamples, numSamples)
	streamer := strmr.NewStreamer(samples)

	c := coordinates.FromGRPC(data.GetCoordinates())
	// TODO: Use this distance to determine volume
	_ = c.Distance(client.coordinates)

	volume := &effects.Volume{
		Streamer: streamer,
		Base:     0,
		Volume:   0,
		Silent:   false,
	}

	return id, volume
}

// SetResampleRate sets the sample rate for audio to be resampled to
func (client *Client) SetResampleRate(sampleRate int) {
	client.resampleRate = beep.SampleRate(sampleRate)
}

// SetResampleRate sets the sample rate for audio to be resampled to
func (client *Client) SetResampleQuality(quality int) {
	client.quality = quality
}

// CloseConn closes conn
func (client *Client) CloseConn() error {
	return client.conn.Close()
}
