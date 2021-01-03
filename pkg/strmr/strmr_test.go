package strmr_test

import (
	"testing"
	"math/rand"
	"time"

	"github.com/aleitner/blather/pkg/strmr"
	"github.com/stretchr/testify/assert"
)

func generateTestSamples(size int) [][2]float64 {
	samples := make([][2]float64, size)

	for i:= 0; i < size; i++ {
		samples[i][0] = rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
		samples[i][1] = rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	}

	return samples
}

func TestStreamer(t *testing.T) {

	{ // perfect scenario
		samples := generateTestSamples(512)

		stream := strmr.NewStreamer(samples)

		numSamplesToRead := 512
		readSamples := make([][2]float64, numSamplesToRead)
		n, ok := stream.Stream(readSamples)

		// We streamed the appropriate number if samples
		assert.Equal(t, 512, n)
		assert.Equal(t, samples[:n], readSamples[:n])
		assert.True(t, ok)

		n, ok = stream.Stream(readSamples)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, readSamples[:n])
		assert.False(t, ok)
	}

	{ // Number of samples to read is greater than the actual number of samples
		samples := generateTestSamples(512)

		stream := strmr.NewStreamer(samples)

		numSamplesToRead := 1024
		actual := make([][2]float64, numSamplesToRead)
		n, ok := stream.Stream(actual)

		// We streamed the appropriate number if samples
		assert.Equal(t, len(samples), n)
		assert.Equal(t, samples[:n], actual[:n])
		assert.True(t, ok)

		n, ok = stream.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
	}

	{ // Number of samples to read is less than the actual number of samples
		samples := generateTestSamples(512)

		stream := strmr.NewStreamer(samples)

		numSamplesToRead := 256
		actual := make([][2]float64, numSamplesToRead)

		totalRead:= 0
		for totalRead < len(samples) {
			n, ok := stream.Stream(actual)
			assert.Equal(t, numSamplesToRead, n)
			assert.Equal(t, samples[totalRead:totalRead+n], actual[:n])
			assert.True(t, ok)
			totalRead += n
		}

		n, ok := stream.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
		assert.Equal(t, len(samples), totalRead)

	}
}
