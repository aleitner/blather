package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/aleitner/spacialPhone/pkg/grpc/client"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"gopkg.in/metakeule/loop.v4"
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
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc Dial fail: %s/n", err)
	}

	var id = rand.Int()
	var logger = log.New()
	client := client.NewContactConnection(id, logger, conn)
	defer client.CloseConn()

	r := loop.New([]byte("Howdy from the client"))

	app.Commands = []cli.Command{
		{
			Name:    "call",
			Aliases: []string{"c"},
			Usage:   "call",
			Action: func(c *cli.Context) error {
				err := client.Call(context.Background(), r)
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
