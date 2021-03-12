package queue_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/aleitner/blather/pkg/queue"
	"github.com/aleitner/blather/pkg/strmr"
	"github.com/stretchr/testify/assert"
)

func generateTestSamples(size int) [][2]float64 {
	samples := make([][2]float64, size)

	for i := 0; i < size; i++ {
		samples[i][0] = rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
		samples[i][1] = rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	}

	return samples
}

func TestQueue(t *testing.T) {

	{ // perfect scenario
		q := queue.NewQueue()

		samples1 := generateTestSamples(512)
		stream := strmr.NewStreamer(samples1)
		q.Add(stream)

		samples2 := generateTestSamples(512)
		stream = strmr.NewStreamer(samples2)
		q.Add(stream)

		actual := make([][2]float64, len(samples2)+len(samples1))

		n, ok := q.Stream(actual)
		assert.Equal(t, len(samples2)+len(samples1), n)

		var samples [][2]float64
		samples = append(samples, samples1...)
		samples = append(samples, samples2...)

		assert.Equal(t, samples, actual[:n])
		assert.True(t, ok)

		n, ok = q.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
	}

	{ // Stream buffer is higher than queue
		q := queue.NewQueue()

		samples1 := generateTestSamples(512)
		stream := strmr.NewStreamer(samples1)
		q.Add(stream)

		samples2 := generateTestSamples(512)
		stream = strmr.NewStreamer(samples2)
		q.Add(stream)

		actual := make([][2]float64, 2048)

		n, ok := q.Stream(actual)
		assert.Equal(t, len(samples2)+len(samples1), n)

		var samples [][2]float64
		samples = append(samples, samples1...)
		samples = append(samples, samples2...)

		assert.Equal(t, samples, actual[:n])
		assert.True(t, ok)

		n, ok = q.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)
	}

	{ // Stream buffer is smaller than queue total
		q := queue.NewQueue()

		samples1 := generateTestSamples(512)
		stream := strmr.NewStreamer(samples1)
		q.Add(stream)

		samples2 := generateTestSamples(512)
		stream = strmr.NewStreamer(samples2)
		q.Add(stream)

		var samples [][2]float64
		samples = append(samples, samples1...)
		samples = append(samples, samples2...)

		lenToRead := 256
		actual := make([][2]float64, lenToRead)

		totalRead := 0
		for totalRead < len(samples) {
			n, ok := q.Stream(actual)
			assert.Equal(t, lenToRead, n)
			assert.Equal(t, samples[totalRead:totalRead+n], actual[:n])
			assert.True(t, ok)
			totalRead += n
		}

		n, ok := q.Stream(actual)
		assert.Equal(t, 0, n)
		assert.Equal(t, [][2]float64{}, actual[:n])
		assert.False(t, ok)

		assert.Equal(t, len(samples), totalRead)
	}
}
