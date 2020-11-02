package muxer

import (
	"sync"

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

func (m *Muxer) Add(id userid.ID, data *call.CallData) {
}

func (m *Muxer) Delete(id userid.ID) {
	m.streamerQueues.Delete(id)
}

func (m *Muxer) Mux() {

}
