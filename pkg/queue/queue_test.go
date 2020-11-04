package queue_test

import (
	"testing"

	"github.com/aleitner/blather/pkg/queue"
	"github.com/aleitner/blather/pkg/strmr"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	{ // perfect scenario
		q := queue.NewQueue()

		samples1 := [][2]float64{[2]float64{1,1},[2]float64{2,2},[2]float64{3,3}}
		stream := strmr.NewStreamer(samples1, len(samples1))
		q.Add(stream)

		samples2 := [][2]float64{[2]float64{4,4},[2]float64{5,5},[2]float64{6,6}}
		stream = strmr.NewStreamer(samples2, len(samples2))
		q.Add(stream)

		actual := make([][2]float64, len(samples2) + len(samples1))

		n, ok := q.Stream(actual)
		assert.Equal(t, len(samples2) + len(samples1), n)

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
}

