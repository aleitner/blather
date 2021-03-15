package muxer_test

import (
	"testing"

	"github.com/aleitner/blather/pkg/muxer"
	"github.com/aleitner/blather/pkg/strmr"
	"github.com/faiface/beep/effects"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMuxer(t *testing.T) {

	{ // Perfect Scenario
		m := muxer.NewMuxer(log.New())
		samples1 := [][2]float64{[2]float64{1, 2}}
		m.Add("1", &effects.Volume{Streamer: strmr.NewStreamer(samples1)})

		samples2 := [][2]float64{[2]float64{3, 4}}
		m.Add("1", &effects.Volume{Streamer: strmr.NewStreamer(samples2)})

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
		m := muxer.NewMuxer(log.New())
		samples1 := [][2]float64{[2]float64{1, 2}}
		m.Add("1", &effects.Volume{Streamer: strmr.NewStreamer(samples1)})

		samples2 := [][2]float64{[2]float64{3, 4}}
		m.Add("2", &effects.Volume{Streamer: strmr.NewStreamer(samples2)})

		expectedSamples := [][2]float64{[2]float64{4, 6}}

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
		m := muxer.NewMuxer(log.New())
		samples1 := [][2]float64{[2]float64{1, 1}, [2]float64{1, 1}}
		m.Add("1", &effects.Volume{Streamer: strmr.NewStreamer(samples1)})

		samples2 := [][2]float64{[2]float64{4, 4}}
		m.Add("2", &effects.Volume{Streamer: strmr.NewStreamer(samples2)})

		expectedSamples := [][2]float64{[2]float64{5, 5}, [2]float64{1, 1}}

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
		m := muxer.NewMuxer(log.New())
		samples1 := [][2]float64{[2]float64{1, 1}, [2]float64{1, 1}, [2]float64{1, 1}, [2]float64{1, 1}}
		m.Add("1", &effects.Volume{Streamer: strmr.NewStreamer(samples1)})

		samples2 := [][2]float64{[2]float64{4, 4}}
		m.Add("2", &effects.Volume{Streamer: strmr.NewStreamer(samples2)})

		expectedSamples := [][2]float64{[2]float64{5, 5}, [2]float64{1, 1}, [2]float64{1, 1}}

		actualSamples := make([][2]float64, 3)
		n, ok := m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, len(expectedSamples), n)
		assert.Equal(t, expectedSamples, actualSamples[:n])

		samples3 := [][2]float64{[2]float64{1, 1}}
		m.Add("2", &effects.Volume{Streamer: strmr.NewStreamer(samples3)})
		n, ok = m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, 1, n)
		assert.Equal(t, [][2]float64{[2]float64{2, 2}}, actualSamples[:n])

		n, ok = m.Stream(actualSamples)
		assert.True(t, ok)
		assert.Equal(t, 0, n)
	}
}
