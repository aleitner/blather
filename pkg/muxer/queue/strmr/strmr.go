package strmr

import "sync"

type Streamer struct {
	mtx        sync.Mutex
	samples    [][2]float64
	numSamples int
}

func NewStreamer(samples [][2]float64, numSamples int) *Streamer {
	if numSamples <= 0 {
		return nil
	}

	if numSamples > len(samples) {
		numSamples = len(samples)
	}

	return &Streamer{
		samples:    samples,
		numSamples: numSamples,
	}
}

func (s *Streamer) Stream(samples [][2]float64) (n int, ok bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	
	// Stream is already empty
	if s.numSamples == 0 {
		return 0, false
	}
	
	numSamplesStreamed := 0
	for i := range samples {
		if s.numSamples == 0 {
			break
		}

		samples[i] = s.samples[i]
		s.numSamples--
		numSamplesStreamed++
	}
	
	return numSamplesStreamed, true
}

func (s *Streamer) Err() error {
	return nil
}
