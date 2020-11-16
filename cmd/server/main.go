package main

import (
	"fmt"
	"net"
	"os"

	"github.com/aleitner/blather/pkg/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	// start grpc server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	logger := log.New()
	s := server.NewBlatherServer(logger)
	server.RegisterBlatherServer(grpcServer, s)

	defer grpcServer.GracefulStop()
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}
}
