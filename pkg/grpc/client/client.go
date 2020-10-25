package client

import (
	"context"
	"fmt"
	"io"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
	"github.com/aleitner/spacialPhone/pkg/user"
	"google.golang.org/grpc"
)

type CallClient interface {
	Connect(ctx context.Context) (*user.User, error)
	Call(ctx context.Context, audioInput io.Reader) error
	CloseConn() error
}

type Client struct {
	route call.CallClient
	conn  *grpc.ClientConn
}

func NewContactConnection(conn *grpc.ClientConn) CallClient {
	return &Client{conn: conn, route: call.NewCallClient(conn)}
}

func (client *Client) Connect(ctx context.Context) (*user.User, error) {
	connectionReq := &call.ConnectRequest{Id: 1}

	_, err := client.route.Connect(ctx, connectionReq)
	if err != nil {
		return nil, fmt.Errorf("Call fail: %s/n", err)
	}

	return nil, nil
}

func (client *Client) Call(ctx context.Context, audioInput io.Reader) error {

	return nil
}

// CloseConn closes conn
func (client *Client) CloseConn() error {
	return client.conn.Close()
}
