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
	logger := log.New()

	// start grpc server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}
	defer lis.Close()

	logger.Infof("Now serving %s", lis.Addr().String())

	grpcServer := grpc.NewServer()

	s := server.NewBlatherServer(logger)
	server.RegisterBlatherServer(grpcServer, s)

	defer grpcServer.GracefulStop()
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
