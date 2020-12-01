package window

import (
	"math"
)

// Cosine is a Cosine window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Cosine_window
type Cosine struct {
}

// NewCosine makes a new Cosine window maker.
func NewCosine() *Cosine {
	return &Cosine{}
}

// Make Cosine window of a given size (number of samples).
func (w *Cosine) Make(size int) []float64 {
	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)

	for n := 0; n < size; n++ {
		coeff[n] = math.Sin(math.Pi * float64(n) / sizeMinus1)
	}

	return coeff
}
