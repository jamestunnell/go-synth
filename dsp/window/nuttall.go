package window

import "math"

// Nuttall is a Nuttall window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Nuttall_window.2C_continuous_first_derivative.
type Nuttall struct {
}

// NewNuttall makes a new Nuttall window maker.
func NewNuttall() *Nuttall {
	return &Nuttall{}
}

// Make Nuttall window of a given size (number of samples).
func (w *Nuttall) Make(size int) []float64 {
	const (
		a0 = 0.355768
		a1 = 0.487396
		a2 = 0.144232
		a3 = 0.012604
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
