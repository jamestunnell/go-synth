package osc

import (
	"math"
)

const (
	fourOverPi           = 4.0 / math.Pi
	negFourOverPiSquared = -4.0 / (math.Pi * math.Pi)
	sineP                = 0.225
)

func sineWave(phase float64) float64 {
	y := fourOverPi*phase + negFourOverPiSquared*phase*math.Abs(phase)
	// for extra precision
	return sineP*(y*math.Abs(y)-y) + y // Q * y + P * y * y.abs
}
