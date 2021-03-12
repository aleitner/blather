package forwarder

import (
	"sync"

	call "github.com/aleitner/blather/pkg/protobuf"
	"github.com/aleitner/blather/pkg/userid"
	log "github.com/sirupsen/logrus"
)

// TransferAgent
type TransferAgent interface {
	Send(*call.CallData) error
}

// Forwarder contains a map of all the transfer agents that data needs to be sent to
type Forwarder struct {
	connections    int
	transferAgents sync.Map
	logger         log.Logger
}

// NewForwarder creates a new Forwarder struct
func NewForwarder() *Forwarder {
	return &Forwarder{}
}

// Forward will forward the data from id
func (f *Forwarder) Forward(data *call.CallData) {
	f.transferAgents.Range(func(key interface{}, value interface{}) bool {
		streamId := key.(userid.ID)
		stream := value.(TransferAgent)
		if streamId == userid.ID(data.UserId) { // Don't need to forward data back to sender
			return true
		}

		if err := stream.Send(data); err != nil {
			f.logger.Error(err)
		}

		return true
	})
}

// ConnectionCount will count number of transferAgents
func (f Forwarder) ConnectionCount() int {
	return f.connections
}

// Add a transferAgent by id
func (f *Forwarder) Add(id userid.ID, stream TransferAgent) {
	if _, loaded := f.transferAgents.LoadOrStore(id, stream); !loaded {
		f.connections += 1
	}
}

// Delete a transferAgent by id
func (f *Forwarder) Delete(id userid.ID) {
	if _, loaded := f.transferAgents.LoadAndDelete(id); loaded {
		f.connections -= 1
	}
}
