package muxer

import (
	"sync"

	"github.com/faiface/beep"

	"github.com/aleitner/spacialPhone/pkg/muxer/queue"

	"github.com/aleitner/spacialPhone/internal/utils"
	"github.com/aleitner/spacialPhone/pkg/muxer/queue/strmr"

	call "github.com/aleitner/spacialPhone/pkg/protobuf"

	"github.com/aleitner/spacialPhone/pkg/userid"
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

	newQ := queue.NewQueue()
	q, _ := m.streamerQueues.LoadOrStore(id, newQ)
	q.(*queue.Queue).Add(streamer)
	m.streamerQueues.Store(id, q)
}

func (m *Muxer) Delete(id userid.ID) {
	m.streamerQueues.Delete(id)
}

func (m *Muxer) Stream(samples [][2]float64) (n int, ok bool) {
	m.streamerQueues.Range(func(key interface{}, value interface{}) bool {
		st := value.(beep.Streamer)
		n, ok = st.Stream(samples)

		return false
	})
	return n, true
}

func (m *Muxer) Err() error {
	return nil
}