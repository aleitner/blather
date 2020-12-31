package forwarder

import (
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
	transferAgents map[userid.ID]TransferAgent
	logger         log.Logger
}

// NewForwarder creates a new Forwarder struct
func NewForwarder() *Forwarder {
	return &Forwarder{
		transferAgents: make(map[userid.ID]TransferAgent),
	}
}

// Forward will forward the data from id
func (f *Forwarder) Forward(id userid.ID, data *call.CallData) {
	recipients := f.transferAgents

	for streamId, stream := range recipients {
		if streamId == id { // Don't need to forward data back to sender
			continue
		}

		if err := stream.Send(data); err != nil {
			f.logger.Error(err)
		}

	}
}

// ConnectionCount will count number of transferAgents
func (f Forwarder) ConnectionCount() int {
	return len(f.transferAgents)
}

// Add a transferAgent by id
func (f *Forwarder) Add(id userid.ID, stream TransferAgent) {
	f.transferAgents[id] = stream
}

// Delete a transferAgent by id
func (f *Forwarder) Delete(id userid.ID) {
	delete(f.transferAgents, id)
}
