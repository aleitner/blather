package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/MarkKremer/microphone"
	"github.com/aleitner/blather/pkg/client"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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
	app := cli.NewApp()

	initializeFlags()

	// Set up connection with rpc server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc Dial fail: %s/n", err)
	}

	var id = rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	var logger = log.New()
	client := client.NewClient(id, logger, conn)
	defer client.CloseConn()

	app.Commands = []cli.Command{
		{
			Name:    "call",
			Aliases: []string{"c"},
			Usage:   "call",
			Action: func(c *cli.Context) error {
				err := microphone.Init()
				if err != nil {
					return err
				}
				defer microphone.Terminate()

				streamer, format, err := microphone.OpenDefaultStream(44100, 2)
				if err != nil {
					return fmt.Errorf("Failed to open default stream: %s", err)
				}

				// Close the stream at the end if it hasn't already been
				// closed explicitly.
				defer streamer.Close()

				err = client.Call(context.Background(), streamer, format)
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
