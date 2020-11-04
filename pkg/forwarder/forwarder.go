package forwarder

import (
	"sync"

	call "github.com/aleitner/spacialPhone/pkg/protobuf"
	"github.com/aleitner/spacialPhone/pkg/userid"
	log "github.com/sirupsen/logrus"
)

// TransferAgent
type TransferAgent interface {
	Send(*call.CallData) error
}

// Forwarder contains a map of all the transfer agents that data needs to be sent to
type Forwarder struct {
	transferAgents sync.Map
	logger         log.Logger
}

// NewForwarder creates a new Forwarder struct
func NewForwarder() *Forwarder {
	return &Forwarder{}
}

// Forward will forward the data from id
func (f *Forwarder) Forward(id userid.ID, data *call.CallData) {
	var wg sync.WaitGroup // NB: We probably don't actually need this wait group

	f.transferAgents.Range(func(key interface{}, value interface{}) bool {
		wg.Add(1)

		go func() {
			defer wg.Done()

			streamId := key.(userid.ID)
			stream := value.(TransferAgent)

			if streamId == id { // Don't need to forward data back to sender
				return
			}

			if err := stream.Send(data); err != nil {
				f.logger.Error(err)
			}
		}()
		return true
	})

	wg.Wait()
}

// ConnectionCount will count number of transferAgents
func (f Forwarder) ConnectionCount() int {
	count := 0

	f.transferAgents.Range(func(key interface{}, value interface{}) bool {
		count++
		return true
	})

	return count
}

// Add a transferAgent by id
func (f *Forwarder) Add(id userid.ID, stream TransferAgent) {
	f.transferAgents.Store(id, stream)
}

// Delete a transferAgent by id
func (f *Forwarder) Delete(id userid.ID) {
	f.transferAgents.Delete(id)
}
