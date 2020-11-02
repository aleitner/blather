package muxer

import (
	"sync"

	"github.com/faiface/beep"

	"github.com/aleitner/spacialPhone/internal/muxer/queue"

	"github.com/aleitner/spacialPhone/internal/muxer/queue/strmr"
	"github.com/aleitner/spacialPhone/internal/utils"

	call "github.com/aleitner/spacialPhone/internal/protobuf"

	"github.com/aleitner/spacialPhone/pkg/user/userid"
	log "github.com/sirupsen/logrus"
)

// Muxer
type Muxer struct {
	streamerQueues sync.Map
	logger         *log.Logger
}

func NewMuxer(logger *log.Logger) *Muxer {
	return &Muxer{
		logger: logger,
	}
}

func (m *Muxer) Add(data *call.CallData) {
	// Todo: We need to also think about storing other data
	userMetaData := data.GetUserMetaData()
	audioData := data.GetAudioData()
	grpcSamples := audioData.GetSamples()
	numSamples := int(audioData.GetNumSamples())
	id := userMetaData.GetId()

	samples := utils.ToSampleRate(grpcSamples, numSamples)
	streamer := strmr.NewStreamer(samples, numSamples)

	newQ := queue.NewQueue(m.logger)
	q, _ := m.streamerQueues.LoadOrStore(id, newQ)
	q.(*queue.Queue).Add(streamer)
	m.streamerQueues.Store(id, q)
}

func (m *Muxer) Delete(id userid.ID) {
	m.streamerQueues.Delete(id)
}

func (m *Muxer) Stream(samples [][2]float64) (n int, ok bool) {
	var tmp [512][2]float64

	for len(samples) > 0 {
		toStream := len(tmp)
		if toStream > len(samples) {
			toStream = len(samples)
		}

		// clear the samples
		for i := range samples[:toStream] {
			samples[i] = [2]float64{}
		}

		snMax := 0 // max number of streamed samples in this iteration
		m.streamerQueues.Range(func(key interface{}, value interface{}) bool {
			st := value.(beep.Streamer)
			// mix the stream
			sn, sok := st.Stream(tmp[:toStream])
			if sn > snMax {
				snMax = sn
			}
			ok = ok || sok

			for i := range tmp[:sn] {
				samples[i][0] += tmp[i][0]
				samples[i][1] += tmp[i][1]
			}

			return true
		})

		n += snMax
		if snMax < len(tmp) {
			break
		}
		samples = samples[snMax:]
	}

	// Stream silence
	if len(samples) == 0 {
		for i := range samples {
			samples[i][0] = 0
			samples[i][1] = 0
		}
	}

	return n, ok
}
