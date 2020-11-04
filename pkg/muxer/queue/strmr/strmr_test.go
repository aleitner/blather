package strmr_test

import (
	"testing"

	"github.com/aleitner/blather/pkg/muxer/queue/strmr"
	"github.com/stretchr/testify/assert"
)

func TestStreamer(t *testing.T) {

	{ // perfect scenario
		samples := [][2]float64{[2]float64{3,3},[2]float64{3,3},[2]float64{3,3}}
		numSamples := len(samples)

		stream := strmr.NewStreamer(samples, numSamples)

		actual := make([][2]float64, numSamples)
		n, ok := stream.Stream(actual)

		// We streamed the appropriate number if samples
		assert.Equal(t, numSamples, n)
		assert.Equal(t, samples, actual[:n])
		assert.True(t, ok)

		n, ok = stream.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
	}

	{ // Number of samples is greater than the actual number of samples
		samples := [][2]float64{[2]float64{3,3}}
		numSamples := 2

		stream := strmr.NewStreamer(samples, numSamples)

		actual := make([][2]float64, numSamples)
		n, ok := stream.Stream(actual)

		// We streamed the appropriate number if samples
		assert.Equal(t, len(samples), n)
		assert.Equal(t, samples, actual[:n])
		assert.True(t, ok)

		n, ok = stream.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
	}

	{ // Number of samples to read is greater than the actual number of samples
		samples := [][2]float64{[2]float64{3,3}}
		numSamples := 1

		stream := strmr.NewStreamer(samples, numSamples)

		actual := make([][2]float64, 3)
		n, ok := stream.Stream(actual)

		// We streamed the appropriate number if samples
		assert.Equal(t, len(samples), n)
		assert.Equal(t, samples, actual[:n])
		assert.True(t, ok)

		n, ok = stream.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
	}

	{ // Number of samples to read is less than the actual number of samples
		samples := [][2]float64{[2]float64{3,3}, {3,3}, {3,3}}

		stream := strmr.NewStreamer(samples, len(samples))

		numSamplesToRead := 2
		actual := make([][2]float64, numSamplesToRead)

		n, ok := stream.Stream(actual)
		assert.Equal(t, numSamplesToRead, n)
		assert.Equal(t, samples[:n], actual[:n])
		assert.True(t, ok)

		n, ok = stream.Stream(actual)
		assert.Equal(t, len(samples) - numSamplesToRead, n)
		assert.Equal(t, samples[:n], actual[:n])
		assert.True(t, ok)

		n, ok = stream.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
	}

}
