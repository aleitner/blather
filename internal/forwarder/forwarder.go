package forwarder

import (
	"fmt"
	"sync"

	call "github.com/aleitner/spacialPhone/internal/protobuf"

	"github.com/aleitner/spacialPhone/pkg/user/user_id"
	log "github.com/sirupsen/logrus"
)

type TransferAgent interface {
	Send(*call.CallData) error
}

// Forwarder
type Forwarder struct {
	transferAgents sync.Map
	logger         log.Logger
}

func NewForwarder() *Forwarder {
	return &Forwarder{}
}

func (f *Forwarder) ForwardTo(id user_id.ID, data *call.CallData) {
	var wg sync.WaitGroup // NB: We probably don't actually want this wait group

	f.transferAgents.Range(func(key interface{}, value interface{}) bool {

		streamId := key.(user_id.ID)
		stream := value.(TransferAgent)
		fmt.Println(streamId)
		if streamId == id {
			return true
		}

		wg.Add(1)
		go func() {
			if err := stream.Send(data); err != nil {
				f.logger.Error(err)
			}
			wg.Done()
		}()

		return true
	})

	wg.Wait()
}

func (f *Forwarder) Add(id user_id.ID, stream TransferAgent) {
	f.transferAgents.Store(id, stream)
}

func (f *Forwarder) Delete(id user_id.ID) {
	f.transferAgents.Delete(id)
}
