package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
	"google.golang.org/grpc"
)

type CallClient interface {
	Call(ctx context.Context, audioInput io.Reader) error
	CloseConn() error
}

type Client struct {
	route  call.PhoneClient
	conn   *grpc.ClientConn
	logger log.Logger
}

func NewContactConnection(conn *grpc.ClientConn) CallClient {
	return &Client{conn: conn, route: call.NewPhoneClient(conn)}
}

func (client *Client) Call(ctx context.Context, audioInput io.Reader) error {
	stream, err := client.route.Call(ctx)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	// Send data
	go func() {
		for {
			buf := make([]byte, 4)
			_, err := audioInput.Read(buf)
			if err != nil {
				// server returns with nil
				if err != io.EOF {
					client.logger.Printf("audio read fail: %s/n", err)
				}
				break
			}

			err = stream.Send(&call.CallData{
				AudioData:     buf,
				AudioEncoding: "idk",
				Length:        int64(len(buf)),
			})
			if err != nil {
				// server returns with nil
				if err != io.EOF {
					client.logger.Printf("stream Send fail: %s/n", err)
				}

				break
			}
		}

		err := stream.CloseSend()
		if err != nil {
			client.logger.Printf("close send fail: %s\n", err)
		}
		wg.Done()
	}()
	wg.Add(1)

	// Receive data
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					log.Fatalf("stream Recv fail: %s/n", err)
				}
				break
			}
			fmt.Printf(string(res.GetAudioData()[:res.GetLength()]))
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
