package strmr

import "sync"

type Streamer struct {
	mtx        sync.Mutex
	samples    [][2]float64
	numSamples int
}

func NewStreamer(samples [][2]float64, numSamples int) *Streamer {
	return &Streamer{
		samples:    samples,
		numSamples: numSamples,
	}
}

func (s *Streamer) Stream(samples [][2]float64) (n int, ok bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	// NB: Maybe we should think of erroring if samples is too small
	if s.numSamples == 0 {
		return 0, false
	}

	n = copy(samples, s.samples)

	s.numSamples -= n
	s.samples = s.samples[n:]

	return n, true
}

func (s *Streamer) Err() error {
	return nil
}
