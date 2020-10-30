package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

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

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var id = r1.Int()
	var logger = log.New()
	client := client.NewContactConnection(id, logger, conn)
	defer client.CloseConn()

	app.Commands = []cli.Command{
		{
			Name:    "call",
			Aliases: []string{"c"},
			Usage:   "call",
			Action: func(c *cli.Context) error {
				txt := c.Args().First()
				if len(txt) == 0 {
					txt = "Howdy from the Client"
				}
				r := loop.New([]byte(txt))
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
