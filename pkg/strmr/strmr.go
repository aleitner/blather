package strmr

type Streamer struct {
	samples [][2]float64
}

func NewStreamer(samples [][2]float64) *Streamer {
	if len(samples) <= 0 {
		return nil
	}

	return &Streamer{
		samples: samples,
	}
}

func (s *Streamer) Stream(samples [][2]float64) (n int, ok bool) {
	// Stream is already empty
	if len(s.samples) == 0 {
		return 0, false
	}

	numSamplesStreamed := 0
	for sample := range samples {
		if numSamplesStreamed >= len(s.samples) {
			break
		}

		samples[sample] = s.samples[sample]
		numSamplesStreamed++
	}

	s.samples = s.samples[numSamplesStreamed:]

	return numSamplesStreamed, true
}

func (s *Streamer) Err() error {
	return nil
}
