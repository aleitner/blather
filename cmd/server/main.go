package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
	"github.com/aleitner/spacialPhone/pkg/grpc/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port int
)

func initializeFlags() {
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()
}

func main() {
	initializeFlags()

	// start grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	logger := log.New()
	s := server.NewCallServer(logger)
	call.RegisterPhoneServer(grpcServer, s)

	defer grpcServer.GracefulStop()
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}
}