package window

import (
	"math"
)

// Blackman is a Blackman window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Blackman_window
type Blackman struct {
}

// NewBlackman makes a new Blackman window maker.
func NewBlackman() *Blackman {
	return &Blackman{}
}

// Make Blackman window of a given size (number of samples).
func (w *Blackman) Make(size int) []float64 {
	const (
		alpha = 0.16
		a0    = (1 - alpha) / 2.0
		a1    = 0.5
		a2    = alpha / 2.0
	)

	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)

	for n := 0; n < size; n++ {
		nFlt := float64(n)

		coeff[n] = a0 -
			(a1 * math.Cos((twoPi*nFlt)/sizeMinus1)) +
			(a2 * math.Cos((fourPi*nFlt)/sizeMinus1))
	}

	return coeff
}
