package muxer_test

import (
	"github.com/aleitner/blather/internal/utils"
	"github.com/aleitner/blather/pkg/muxer"
	blatherpb "github.com/aleitner/blather/pkg/protobuf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateCallData(samples [][2]float64, sampleNum, sampleRate, id int) *blatherpb.CallData {
	utils.ToGRPCSampleRate(samples, sampleNum)
	return &blatherpb.CallData{
		AudioData:    &blatherpb.AudioData{
			AudioEncoding: "test",
			Samples:       utils.ToGRPCSampleRate(samples, sampleNum),
			NumSamples:    uint32(sampleNum),
			Format:        &blatherpb.Format{
				SampleRate:  uint32(sampleRate),
				NumChannels: 2,
				Precision:   0,
			},
		},
		UserMetaData: &blatherpb.UserMetaData{
			Id:          uint64(id),
			Coordinates: &blatherpb.Coordinates{
				X: 0,
				Y: 0,
				Z: 0,
			},
		},
	}
}

func TestMuxer(t *testing.T) {

	{ // Perfect Scenario
		m := muxer.NewMuxer()
		samples1 := [][2]float64{[2]float64{1,2}}
		data1 := generateCallData(samples1, len(samples1), 44100, 1)
		m.Add(data1)

		samples2 := [][2]float64{[2]float64{3,4}}
		data2 := generateCallData(samples2, len(samples2), 44100, 1)
		m.Add(data2)

		var expectedSamples [][2]float64
		expectedSamples = append(expectedSamples, samples1...)
		expectedSamples = append(expectedSamples, samples2...)

		actualSamples := make([][2]float64, 10)
		n, ok := m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, len(expectedSamples), n)
		assert.Equal(t, expectedSamples, actualSamples[:n])

		n, ok = m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, 0, n)
	}

	{ // Properly Mix Audio
		m := muxer.NewMuxer()
		samples1 := [][2]float64{[2]float64{1,2}}
		data1 := generateCallData(samples1, len(samples1), 44100, 1)
		m.Add(data1)

		samples2 := [][2]float64{[2]float64{3,4}}
		data2 := generateCallData(samples2, len(samples2), 44100, 2)
		m.Add(data2)

		expectedSamples := [][2]float64{[2]float64{4,6}}

		actualSamples := make([][2]float64, 10)
		n, ok := m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, len(expectedSamples), n)
		assert.Equal(t, expectedSamples, actualSamples[:n])

		n, ok = m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, 0, n)
	}

	{ // Properly Mix Audio of varying lengths
		m := muxer.NewMuxer()
		samples1 := [][2]float64{[2]float64{1,1},[2]float64{1,1}}
		data1 := generateCallData(samples1, len(samples1), 44100, 1)
		m.Add(data1)

		samples2 := [][2]float64{[2]float64{4,4}}
		data2 := generateCallData(samples2, len(samples2), 44100, 2)
		m.Add(data2)

		expectedSamples := [][2]float64{[2]float64{5,5}, [2]float64{1,1}}

		actualSamples := make([][2]float64, 10)
		n, ok := m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, len(expectedSamples), n)
		assert.Equal(t, expectedSamples, actualSamples[:n])

		n, ok = m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, 0, n)
	}
}