package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/aleitner/spacialPhone/pkg/grpc/client"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"strings"
)

var (
	port int
)

func initializeFlags() {
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()
}

func main() {
	app := cli.NewApp()

	initializeFlags()

	// Set up connection with rpc server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc Dial fail: %s/n", err)
	}

	client := client.NewContactConnection(conn)
	defer client.CloseConn()

	r:= strings.NewReader("Howdy from the client")

	app.Commands = []cli.Command{
		{
			Name:    "call",
			Aliases: []string{"c"},
			Usage:   "call",
			Action: func(c *cli.Context) error {
				err := client.Call(context.Background(),r)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("app Run fail: %s/n", err)
	}
}