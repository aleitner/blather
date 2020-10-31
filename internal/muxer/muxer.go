package muxer

import (
	"bytes"
	"sync"

	"github.com/aleitner/spacialPhone/pkg/user/user_id"
	log "github.com/sirupsen/logrus"
)

// Muxer
type Muxer struct {
	readers sync.Map
	logger  *log.Logger
}

func NewMuxer(logger *log.Logger) *Muxer {
	return &Muxer{
		logger: logger,
	}
}

func (m *Muxer) Add(id user_id.ID, data []byte) {
	buffer := bytes.NewBuffer(data)
	loadedBuffer, loaded := m.readers.LoadOrStore(id, buffer)
	if !loaded {
		return
	}

	if _, err := loadedBuffer.(*bytes.Buffer).Write(data); err != nil {
		m.logger.Errorf("Failed to Add data to muxer. %s", err.Error())
	}
}

func (m *Muxer) Delete(id user_id.ID) {
	m.readers.Delete(id)
}

func (m *Muxer) Mux() {

}
