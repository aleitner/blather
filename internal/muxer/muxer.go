package muxer

import "sync"

type Muxer struct {
	mtx     sync.Mutex
	sources map[int][]byte
}

func (m *Muxer) Add(id int) {

}

func (m *Muxer) Delete(id int) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.sources, id)
}

func (m *Muxer) Mux(id int) {
}
