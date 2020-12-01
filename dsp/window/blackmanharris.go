package window

import (
	"math"
)

// BlackmanHarris is a BlackmanHarris window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Blackman-Harris_window
type BlackmanHarris struct {
}

// NewBlackmanHarris makes a new BlackmanHarris window maker.
func NewBlackmanHarris() *BlackmanHarris {
	return &BlackmanHarris{}
}

// Make BlackmanHarris window of a given size (number of samples).
func (w *BlackmanHarris) Make(size int) []float64 {
	const (
		a0 = 0.35875
		a1 = 0.48829
		a2 = 0.14128
		a3 = 0.01168
	)

	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)

	for n := 0; n < size; n++ {
		nFlt := float64(n)

		coeff[n] = a0 -
			a1*math.Cos((twoPi*nFlt)/sizeMinus1) +
			a2*math.Cos((fourPi*nFlt)/sizeMinus1) -
			a3*math.Cos((sixPi*nFlt)/sizeMinus1)
	}

	return coeff
}
