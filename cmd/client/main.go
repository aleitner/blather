package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aleitner/blather/pkg/client"
	"github.com/faiface/beep/mp3"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "address",
				Value: ":8080",
				Usage: "port of server being ",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "create a call room",
				Action: func(c *cli.Context) error {
					// Set up connection with rpc server
					var conn *grpc.ClientConn
					conn, err := grpc.Dial(c.String("address"), grpc.WithInsecure())
					if err != nil {
						log.Fatalf("grpc Dial fail: %s/n", err)
					}

					var id = rand.New(rand.NewSource(time.Now().UnixNano())).Int()

					var logger = log.New()
					client := client.NewClient(id, logger, conn)
					defer client.CloseConn()

					ctx:= context.Background()

					roomID, err := client.CreateRoom(ctx)
					if err != nil {
						return err
					}

					fmt.Printf("ROOM: %s\n", roomID)
					return nil
				},
			},
			{
				Name:    "join",
				Aliases: []string{"j"},
				Usage:   "join a call room",
				Flags: []cli.Flag {
					&cli.StringFlag{
						Name: "room",
						Value: "",
						Usage: "room number",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					// Set up connection with rpc server
					var conn *grpc.ClientConn
					conn, err := grpc.Dial(c.String("address"), grpc.WithInsecure())
					if err != nil {
						log.Fatalf("grpc Dial fail: %s/n", err)
					}

					var id = rand.New(rand.NewSource(time.Now().UnixNano())).Int()

					var logger = log.New()
					client := client.NewClient(id, logger, conn)
					defer client.CloseConn()

					filepath := c.Args().First()
					f, err := os.Open(filepath)
					if err != nil {
						return err
					}

					ctx:= context.Background()

					streamer, format, err := mp3.Decode(f)
					if err != nil {
						return err
					}
					defer streamer.Close()

					err = client.Call(ctx, c.String("room"), streamer, format)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
