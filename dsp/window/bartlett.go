package window

import (
	"math"
)

// Bartlett is a Bartlett window maker.
// Bartlett window is very similar to triangular window.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Triangular_window
type Bartlett struct {
}

// NewBartlett makes a new Bartlett window maker.
func NewBartlett() *Bartlett {
	return &Bartlett{}
}

// Make Bartlett window of a given size (number of samples).
func (w *Bartlett) Make(size int) []float64 {
	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)
	sizeMinus1Over2 := float64(size-1) / 2.0

	for n := 0; n < size; n++ {
		x := sizeMinus1Over2 - math.Abs(float64(n)-sizeMinus1Over2)

		coeff[n] = (2.0 / sizeMinus1) * x
	}

	return coeff
}
