package osc

import (
	"math"
)

const (
	oneOverPi = 1.0 / math.Pi
)

type Sawtooth struct {
	*Osc
}

// NewSawtooth makes a sine wave oscillator.
func NewSawtooth() *Sawtooth {
	return &Sawtooth{
		Osc: New(sawtoothWave),
	}
}

func sawtoothWave(phase float64) float64 {
	return phase * oneOverPi
}
