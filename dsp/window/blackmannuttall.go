package window

import (
	"math"
)

// BlackmanNuttall is a BlackmanNuttall window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Blackman-Nuttall_window
type BlackmanNuttall struct {
}

// NewBlackmanNuttall makes a new BlackmanNuttall window maker.
func NewBlackmanNuttall() *BlackmanNuttall {
	return &BlackmanNuttall{}
}

// Make BlackmanNuttall window of a given size (number of samples).
func (w *BlackmanNuttall) Make(size int) []float64 {
	const (
		a0 = 0.3635819
		a1 = 0.4891775
		a2 = 0.1365995
		a3 = 0.0106411
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
