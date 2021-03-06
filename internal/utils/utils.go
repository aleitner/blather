package utils

import (
	"log"
	"math/rand"
	"net"

	call "github.com/aleitner/blather/pkg/protobuf"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func ToGRPCSampleRate(samples [][2]float64, numSamples int) []*call.Sample {
	grpcSamples := make([]*call.Sample, numSamples)
	for i := 0; i < numSamples; i++ {
		grpcSamples[i] = &call.Sample{
			LeftChannel:  samples[i][0],
			RightChannel: samples[i][1],
		}
	}

	return grpcSamples
}

func ToSampleRate(grpcSamples []*call.Sample, numSamples int) [][2]float64 {
	samples := make([][2]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		samples[i][0] = grpcSamples[i].GetLeftChannel()
		samples[i][1] = grpcSamples[i].GetRightChannel()
	}

	return samples
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
