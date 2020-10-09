package osc

import "math"

const (
	twoOverPi = 2.0 / math.Pi
)

func Triangle(params *Params) *Osc {
	return new(params, triangleWave)
}

func triangleWave(phase float64) float64 {
	return math.Abs(twoOverPi*phase) - 1.0
}
