package window

import (
	"math"
)

// BartlettHann is a BartlettHann window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Bartlett-Hann_window
type BartlettHann struct {
}

// NewBartlettHann makes a new BartlettHann window maker.
func NewBartlettHann() *BartlettHann {
	return &BartlettHann{}
}

// Make BartlettHann window of a given size (number of samples).
func (w *BartlettHann) Make(size int) []float64 {
	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)
	sizeMinus1Over2 := float64(size-1) / 2.0

	for n := 0; n < size; n++ {
		coeff[n] = (2.0 / sizeMinus1) * (sizeMinus1Over2 - math.Abs(float64(n)-sizeMinus1Over2))
	}

	return coeff
}
