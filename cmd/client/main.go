package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aleitner/blather/pkg/client"
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
					client := client.NewClient(id, logger, conn)
					defer client.CloseConn()

					ctx := context.Background()

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

					mctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
						fmt.Printf("LOG <%v>\n", message)
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
					deviceConfig.SampleRate = 22050

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

					var logger = log.New()
					client := client.NewClient(id, logger, conn)
					defer client.CloseConn()

					ctrlc := make(chan os.Signal)
					signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)
					go func() {
						<-ctrlc
						fmt.Println("\r- Turning off microphone...")
						stream.Close()
						client.CloseConn()
						os.Exit(0)
					}()

					speaker.Play(beep.Seq(client.Muxer))

					err = client.Call(ctx, c.String("room"), stream, format)
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
