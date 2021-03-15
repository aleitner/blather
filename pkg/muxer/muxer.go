package muxer

import (
	"sync"

	"github.com/faiface/beep"
	log "github.com/sirupsen/logrus"

	"github.com/aleitner/blather/pkg/queue"
	"github.com/aleitner/blather/pkg/userid"
)

// Muxer
type Muxer struct {
	logger *log.Logger
	Queues sync.Map
	numQueues int
}

func NewMuxer(logger *log.Logger) *Muxer {
	return &Muxer{
		logger: logger,
	}
}

func (m *Muxer) Add(streamerID userid.ID, streamer beep.Streamer) {
	q, loaded := m.Queues.LoadOrStore(streamerID, queue.NewQueue())

	if !loaded {
		m.numQueues += 1
	}

	q.(*queue.Queue).Add(streamer)
}

func (m *Muxer) Delete(id userid.ID) {
	if _, loaded := m.Queues.LoadAndDelete(id); loaded {
		m.numQueues -= 1
	}
}

func (m *Muxer) Len() int {
	return m.numQueues
}

func (m *Muxer) Stream(samples [][2]float64) (n int, ok bool) {
	streamedCount := make(map[userid.ID]int, m.Len())
	var tmp [512][2]float64

	toStream := len(tmp)
	if toStream > len(samples) {
		toStream = len(samples)
	}

	// clear the samples
	for i := range samples[:toStream] {
		samples[i] = [2]float64{}
	}

	n = 0

	for m.Len() > 0 && n < toStream {
		m.Queues.Range(func(key interface{}, value interface{}) bool {
		id := key.(userid.ID)
		st := value.(*queue.Queue)

			_, bok := streamedCount[id]
			if !bok {
				streamedCount[id] = 0
			}
			// mix the stream
			sn, sok := st.Stream(tmp[streamedCount[id]:toStream])

			for i := range tmp[:sn] {
				samples[i][0] += tmp[i][0]
				samples[i][1] += tmp[i][1]
			}

			streamedCount[id] += sn

			if streamedCount[id] > n {
				n = streamedCount[id]
			}

			if !sok {
				m.Delete(id)
			}

			return true
		})
	}
	return n, true
}

func (m *Muxer) Err() error {
	return nil
}
