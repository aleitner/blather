package server

import (
	"context"
	call "github.com/aleitner/spacialPhone/internal/protobuf"
)

type CallServer struct {
}

func (cs *CallServer) Connect(ctx context.Context, req *call.ConnectRequest) (*call.ConnectResponse, error) {
	return nil, nil
}

func (cs *CallServer) Call(stream call.Call_CallServer) error {
	return nil
}
