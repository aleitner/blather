package forwarder_test

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/aleitner/blather/pkg/forwarder"
	blatherpb "github.com/aleitner/blather/pkg/protobuf"
	"github.com/aleitner/blather/pkg/userid"
	"github.com/stretchr/testify/require"
)

type MockTransferAgent struct {
	store bool
	ReceivedFromCount map[uint64]int
}

func NewMockTransferAgent(store bool) *MockTransferAgent {
	return &MockTransferAgent{
		store: store,
		ReceivedFromCount: make(map[uint64]int),
	}
}

func (mta *MockTransferAgent) Send(data *blatherpb.CallData) error {
	if !mta.store {
		return nil
	}

	mta.ReceivedFromCount[data.UserId] += 1
	return nil
}

func TestForwarder(t *testing.T) {

	{ // Handling multiple connections
		f := forwarder.NewForwarder()
		require.Equal(t, 0, f.ConnectionCount())

		mta1 := NewMockTransferAgent(true)
		mta2 := NewMockTransferAgent(true)
		mta3 := NewMockTransferAgent(true)

		// Add a transfer agents by their user id
		f.Add(userid.ID(1), mta1)
		f.Add(userid.ID(2), mta2)
		f.Add(userid.ID(3), mta3)
		require.Equal(t, 3, f.ConnectionCount())

		// forward the calldata
		f.Forward(&blatherpb.CallData{
			UserId:    1,
		})

		require.Equal(t, 0, mta1.ReceivedFromCount[1])
		require.Equal(t, 1, mta2.ReceivedFromCount[1])
		require.Equal(t, 1, mta3.ReceivedFromCount[1])
	}

	{ // Handling 1 connection
		f := forwarder.NewForwarder()
		require.Equal(t, 0, f.ConnectionCount())

		mta := NewMockTransferAgent(true)

		// Add a transfer agent by their user id
		f.Add(userid.ID(1), mta)
		require.Equal(t, 1, f.ConnectionCount())

		// Don't add the same user ID twice
		f.Add(userid.ID(1), mta)
		require.Equal(t, 1, f.ConnectionCount())

		calldata1 := &blatherpb.CallData{
			UserId:    1,
		}

		// forward the calldata
		f.Forward(calldata1)
		require.Equal(t, 0, mta.ReceivedFromCount[1])

		// delete transfer agent by userid
		f.Delete(userid.ID(1))
		require.Equal(t, 0, f.ConnectionCount())

		// don't lower connection count if userid doesn't exist
		f.Delete(userid.ID(1))
		require.Equal(t, 0, f.ConnectionCount())
	}

	// TODO: We need to think of a cleaner way to check for speed and concurrent transfer failure
	// This test is for how quickly data transfers
	{ // Delete users while transferring data
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		f := forwarder.NewForwarder()

		// Prepare the transfer agents
		transferAgentCount := 1000
		for i := 0; i < transferAgentCount; i++ {
			mta := NewMockTransferAgent(false)
			f.Add(userid.ID(i), mta)
		}

		// Periodically delete
		// Delete any of the forwarders except the last 10
		go func() {
			x1 := rand.NewSource(time.Now().UnixNano())
			y1 := rand.New(x1)
			ticker := time.NewTicker(time.Millisecond)
			for {
				select {
				case <- ticker.C:
					f.Delete(userid.ID(y1.Intn(transferAgentCount - 10)))
				case <-ctx.Done():
					break
				}
			}
		}()

		// Forward loop
		// Forward data from only the last 10 forwarders
		var wg sync.WaitGroup
		for i := transferAgentCount - 10; i < transferAgentCount; i++{
			wg.Add(1)
			go func(i int) {
				for x := 0; x < 1000; x++ {
					f.Forward(&blatherpb.CallData{
						UserId: uint64(i),
					})
				}
				wg.Done()
			}(i)
		}

		wg.Wait()
	}

}