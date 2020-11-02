package utils

import (
	"log"
	"net"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
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
