package client

import (
	"context"
	"io"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
	"google.golang.org/grpc"
)

type CallClient interface {
	Call(ctx context.Context, audioInput io.Reader) error
	CloseConn() error
}

type Client struct {
	route call.PhoneClient
	conn  *grpc.ClientConn
}

func NewContactConnection(conn *grpc.ClientConn) CallClient {
	return &Client{conn: conn, route: call.NewPhoneClient(conn)}
}

func (client *Client) Call(ctx context.Context, audioInput io.Reader) error {

	return nil
}

// CloseConn closes conn
func (client *Client) CloseConn() error {
	return client.conn.Close()
}
