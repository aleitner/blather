package queue

import (
	"sync"

	"github.com/aleitner/blather/pkg/muxer/queue/strmr"
)

type Queue struct {
	mtx       sync.Mutex
	streamers []*strmr.Streamer
}

func NewQueue() *Queue {
	return &Queue{
		streamers: make([]*strmr.Streamer, 0),
	}
}

func (q *Queue) Add(streamer *strmr.Streamer) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.streamers = append(q.streamers, streamer)
}

func (q *Queue) Stream(samples [][2]float64) (n int, ok bool) {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	filled := 0
	for filled < len(samples) {
		if len(q.streamers) == 0 {
			break
		}

		toStream := len(samples) - filled
		buf := make([][2]float64, toStream)

		// NB: See if we can find a way to write directly into buf at different points
		// We are currently making two copies of the same thing which is not performant
		n, ok = q.streamers[0].Stream(buf)
		if !ok {
			q.streamers = q.streamers[1:]
		}

		for i := 0; i < n; i++ {
			samples[i+filled] = buf[i]
		}

		filled += n
	}

	return filled, filled > 0
}

func (q *Queue) Err() error {
	return nil
}
