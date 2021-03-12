package queue

import (
	"math"

	"github.com/aleitner/blather/pkg/strmr"
)

// Volume adjusts the volume of the wrapped Streamer in a human-natural way. Human's perception of
// volume is roughly logarithmic to gain and thus the natural way to adjust volume is exponential.
//
// Natural Base for the exponentiation is somewhere around 2. In order to adjust volume along
// decibells, pick 10 as the Base and set Volume to dB/10. However, adjusting volume along decibells
// is nowhere as natural as with bases around 2.
//
// Volume of 0 means no change. Negative Volume will decrease the perceived volume and positive will
// increase it.
//
// With exponential gain it's impossible to achieve the zero volume. When Silent field is set to
// true, the output is muted.
type Volume struct {
	streamer *Queue
	Base     float64
	Volume   float64
	Silent   bool
}

// NewVolume creates a new Volume and Streamer Queue
func NewVolume() *Volume {
	return &Volume{
		streamer: NewQueue(),
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
}

// Add streamer to Volume Streamer queue
func (v *Volume) Add(streamer *strmr.Streamer) {
	v.streamer.Add(streamer)
}

// Stream streams the wrapped Streamer with volume adjusted according to Base, Volume and Silent
// fields.
func (v *Volume) Stream(samples [][2]float64) (n int, ok bool) {
	n, ok = v.streamer.Stream(samples)
	gain := 0.0
	if !v.Silent {
		gain = math.Pow(v.Base, v.Volume)
	}
	for i := range samples[:n] {
		samples[i][0] *= gain
		samples[i][1] *= gain
	}
	return n, ok
}

// Err propagates the wrapped Streamer's errors.
func (v *Volume) Err() error {
	return v.streamer.Err()
}

