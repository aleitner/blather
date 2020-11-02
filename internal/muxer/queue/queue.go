package queue

import (
	"sync"

	"github.com/aleitner/spacialPhone/internal/muxer/queue/strmr"
	log "github.com/sirupsen/logrus"
)

type Queue struct {
	logger    *log.Logger
	mtx       sync.Mutex
	streamers []*strmr.Streamer
}

func NewQueue(logger *log.Logger) *Queue {
	return &Queue{
		logger:    logger,
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
	// We use the filled variable to track how many samples we've
	// successfully filled already. We loop until all samples are filled.
	filled := 0
	for filled < len(samples) {
		// There are no streamers in the queue, so we stream silence.
		if len(q.streamers) == 0 {
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		// We stream from the first streamer in the queue.
		n, ok := q.streamers[0].Stream(samples[filled:])
		// If it's drained, we pop it from the queue, thus continuing with
		// the next streamer.
		if !ok {
			q.streamers = q.streamers[1:]
		}
		// We update the number of filled samples.
		filled += n
	}
	return len(samples), true
}

func (q *Queue) Err() error {
	return nil
}
