package osc

import (
	"math"
)

const (
	twoOverPi = 2.0 / math.Pi
)

func triangleWave(phase float64) float64 {
	return math.Abs(twoOverPi*phase) - 1.0
}
