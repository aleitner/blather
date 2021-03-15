package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aleitner/blather/pkg/client"
	"github.com/aleitner/blather/pkg/userid"
	"github.com/aleitner/microphone"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/gen2brain/malgo"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "address",
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
					blatherClient := client.NewClient(userid.FromInt(id), logger, conn)
					defer blatherClient.CloseConn()

					ctx := context.Background()

					roomID, err := blatherClient.CreateRoom(ctx)
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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "room",
						Value:    "",
						Usage:    "room number",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					ctx := context.Background()
					var logger = log.New()

					mctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
						log.Debugf("malgo log: %v\n", message)
					})
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					defer func() {
						_ = mctx.Uninit()
						mctx.Free()
					}()

					deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
					deviceConfig.Capture.Format = malgo.FormatS24
					deviceConfig.Capture.Channels = 2
					deviceConfig.SampleRate = 44100

					stream, format, err := microphone.OpenStream(mctx, deviceConfig)
					if err != nil {
						log.Fatal(err)
					}

					speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

					stream.Start()

					// Set up connection with rpc server
					var conn *grpc.ClientConn
					conn, err = grpc.Dial(c.String("address"), grpc.WithInsecure())
					if err != nil {
						log.Fatalf("grpc Dial fail: %s/n", err)
					}

					var id = rand.New(rand.NewSource(time.Now().UnixNano())).Int()

					blatherClient := client.NewClient(userid.FromInt(id), logger, conn)
					defer blatherClient.CloseConn()

					speaker.Play(beep.Seq(blatherClient.Muxer))

					err = blatherClient.Call(ctx, c.String("room"), stream, int(format.SampleRate))
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
