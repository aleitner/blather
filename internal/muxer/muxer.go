package muxer

import (
	"io"
	"strconv"
	"sync"
)

type ID int

func (id ID) toInt() int {
	return int(id)
}

func (id ID) String() string {
	return strconv.Itoa(id.toInt())
}

func NewID(id string) (ID, error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, nil
	}

	return ID(n), nil
}

// Producer is an reader for data produced from a specific node
type Producer struct {
	mtx  sync.Mutex
	id   ID
	data []byte
}

func NewProducer(id ID) *Producer {
	return &Producer{
		id: id,
	}
}

func (s *Producer) Read(buf []byte) (int, error) {

	// Determine length to be read
	lenToCopy := cap(buf)
	if lenToCopy > 64*1024 {
		lenToCopy = 64 * 1024 // 64 KiB is recommended grpc message size
	} else if lenToCopy <= 0 {
		return 0, nil
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	// NB: Maybe send a message to a cond that waits for len(s.data) to not be 0
	if len(s.data) == 0 {
		return 0, nil
	}

	copy(buf, s.data[:lenToCopy])
	s.data = s.data[:lenToCopy-1]
	return len(s.data), nil
}

// Consumer contains a map of Producers that can be accessed by Producer ID
type Consumer struct {
	id        ID
	producers map[ID]*Producer // ID of producer
}

func NewConsumer(consumerId ID, listOfChatters []ID) *Consumer {
	producers := make(map[ID]*Producer)
	for _, chatterID := range listOfChatters {
		producers[chatterID] = &Producer{
			id: chatterID,
		}
	}

	return &Consumer{
		id:        consumerId,
		producers: producers,
	}
}

// Muxer contains a Consumers that can be accessed by Consumer ID
type Muxer struct {
	consumers map[ID]*Consumer // ID of consumer
}

func NewMuxer() *Muxer {
	return &Muxer{
		consumers: make(map[ID]*Consumer),
	}
}

func (m *Muxer) Produce(id ID, data []byte) {
	var wg sync.WaitGroup // We might not need to actually wait if we add sync cond
	for consumerId, consumer := range m.consumers {
		if consumerId == id {
			continue // No need to save audio for self
		}

		go func() {
			consumer.producers[id].mtx.Lock()
			defer consumer.producers[id].mtx.Unlock()
			consumer.producers[id].data = append(consumer.producers[id].data, data...)
			wg.Done()
		}()
		wg.Add(1)
	}

	wg.Wait()
}

func (m *Muxer) Consume(id ID) []byte {
	var readers []io.Reader
	for producerId, producer := range m.consumers[id].producers {
		if producerId == id {
			continue // No need to consume audio produced by self
		}
		readers = append(readers, producer)
	}

	buf := make([]byte, 64*1024)

	for _, r := range readers {
		r.Read(buf)
	}
	// NB: we might consume too much...
	// TODO: Mix readers

	return buf
}

func (m Muxer) ListOfChatters() (ids []ID) {
	for id, _ := range m.consumers {
		ids = append(ids, id)
	}

	return ids
}

func (m *Muxer) Add(id ID) {
	if len(m.consumers) > 0 {
		for _, consumer := range m.consumers {
			consumer.producers[id] = NewProducer(id)
		}
	}

	m.consumers[id] = NewConsumer(id, m.ListOfChatters())
}

func (m *Muxer) Delete(id ID) {
	// Delete the Producer
	delete(m.consumers, id)

	// Delete all the consumers
	for _, consumer := range m.consumers {
		delete(consumer.producers, id)
	}
}
