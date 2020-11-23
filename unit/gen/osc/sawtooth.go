package osc

import (
	"math"
)

const (
	oneOverPi = 1.0 / math.Pi
)

func sawtoothWave(phase float64) float64 {
	return phase * oneOverPi
}
