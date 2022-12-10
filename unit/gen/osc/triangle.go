package osc

import (
	"math"
)

const (
	twoOverPi = 2.0 / math.Pi
)

type Triangle struct {
	*Osc
}

// NewTriangle makes a triangle wave oscillator.
func NewTriangle() *Triangle {
	return &Triangle{
		Osc: New(triangleWave),
	}
}

func triangleWave(phase float64) float64 {
	return math.Abs(twoOverPi*phase) - 1.0
}
