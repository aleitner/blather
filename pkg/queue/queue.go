package queue

import (
	"github.com/aleitner/blather/pkg/strmr"
)

type Queue struct {
	streamers []*strmr.Streamer
}

func NewQueue() *Queue {
	return &Queue{
		streamers: []*strmr.Streamer{},
	}
}

func (q *Queue) Add(streamer *strmr.Streamer) {
	q.streamers = append(q.streamers, streamer)
}

func (q *Queue) Stream(samples [][2]float64) (filled int, ok bool) {
	for filled < len(samples) {
		if len(q.streamers) == 0 {
			break
		}

		n, ok := q.streamers[0].Stream(samples[filled:])
		if !ok {
			q.streamers = q.streamers[1:]
		}

		filled += n
	}

	return filled, filled > 0
}

func (q *Queue) Err() error {
	return nil
}
