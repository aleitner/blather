package muxer

import (
	log "github.com/sirupsen/logrus"

	"github.com/aleitner/blather/internal/utils"
	blatherpb "github.com/aleitner/blather/pkg/protobuf"
	"github.com/aleitner/blather/pkg/queue"
	"github.com/aleitner/blather/pkg/strmr"
	"github.com/aleitner/blather/pkg/userid"
)

// Muxer
type Muxer struct {
	logger *log.Logger
	Queues map[userid.ID]*queue.Queue
}

func NewMuxer(logger *log.Logger) *Muxer {
	return &Muxer{
		logger: logger,
		Queues: make(map[userid.ID]*queue.Queue),
	}
}

func (m Muxer) Len() int {
	return len(m.Queues)
}

func (m *Muxer) Add(data *blatherpb.CallData) {
	audioData := data.GetAudioData()
	if audioData.GetNumSamples() == 0 {
		return
	}

	grpcSamples := audioData.GetSamples()
	numSamples := int(audioData.GetNumSamples())
	id := userid.ID(data.GetUserId())

	samples := utils.ToSampleRate(grpcSamples, numSamples)
	streamer := strmr.NewStreamer(samples, numSamples)

	q, ok := m.Queues[id]
	if !ok {
		q = queue.NewQueue()
		m.Queues[id] = q
	}
	q.Add(streamer)
}

func (m *Muxer) Delete(id userid.ID) {
	delete(m.Queues, id)
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
			for id, st := range m.Queues {

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
		}
	}
	return n, true
}

func (m *Muxer) Err() error {
	return nil
}
