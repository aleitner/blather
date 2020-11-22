package muxer

import (
	"sync"

	"github.com/aleitner/blather/internal/utils"
	"github.com/aleitner/blather/pkg/protobuf"
	"github.com/aleitner/blather/pkg/queue"
	"github.com/aleitner/blather/pkg/strmr"
	"github.com/aleitner/blather/pkg/userid"
	"github.com/faiface/beep"
)

// Muxer
type Muxer struct {
	Queues sync.Map

	mtx sync.Mutex
	streamerCount int
}

func NewMuxer() *Muxer {
	return &Muxer{
	}
}

func (m Muxer) Len() int {
	return m.streamerCount
}

func (m *Muxer) Add(data *blatherpb.CallData) {
	audioData := data.GetAudioData()
	if audioData.GetNumSamples() == 0 {
		return
	}

	grpcSamples := audioData.GetSamples()
	numSamples := int(audioData.GetNumSamples())
	id := data.GetUserId()

	samples := utils.ToSampleRate(grpcSamples, numSamples)
	streamer := strmr.NewStreamer(samples, numSamples)

	newQ := queue.NewQueue()
	q, _ := m.Queues.LoadOrStore(id, newQ)
	q.(*queue.Queue).Add(streamer)
	m.Queues.Store(id, q)

	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.streamerCount++
}

func (m *Muxer) Delete(id userid.ID) {
	m.Queues.Delete(id)

	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.streamerCount--
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

	for m.Len() > 0 && n < toStream {
		m.Queues.Range(func(key interface{}, value interface{}) bool {
			st := value.(beep.Streamer)
			id := userid.ID(key.(uint64))
			// mix the stream
			sn, sok := st.Stream(tmp[streamedCount[id]:toStream])
			for i := range tmp[:sn] {
				samples[i][0] += tmp[i][0]
				samples[i][1] += tmp[i][1]
			}

			if !sok {
				// remove drained streamer
				m.Delete(id)
			}

			streamedCount[id] += sn

			if streamedCount[id] > n {
				n = streamedCount[id]
			}
			return true
		})
	}

	return n, true
}

func (m *Muxer) Err() error {
	return nil
}