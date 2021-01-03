package strmr

type Streamer struct {
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
	// Stream is already empty
	if s.numSamples == 0 {
		return 0, false
	}

	numSamplesStreamed := 0

	maxsamples := len(samples)
	maxssamplesoffset := len(s.samples) - s.numSamples
	for i := 0; i < maxsamples; i++ {
		if s.numSamples == 0 {
			break
		}
		samples[numSamplesStreamed] = s.samples[maxssamplesoffset+numSamplesStreamed]
		s.numSamples--
		numSamplesStreamed++
	}

	return numSamplesStreamed, true
}

func (s *Streamer) Err() error {
	return nil
}
