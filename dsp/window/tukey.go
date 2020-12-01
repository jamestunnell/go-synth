package window

import (
	"fmt"
	"math"
)

// Tukey is a Tukey window maker.
// The Tukey window, also known as tapered cosine, can be regarded as a cosine
// lobe of width alpha * N / 2 that is convolved with a rectangular window. At
// alpha = 0 it becomes rectangular, and at alpha = 1 it becomes a Hann window.
// For more info, see https://en.wikipedia.org/wiki/Window_function//Tukey_window.
type Tukey struct {
	alpha float64
}

// NewTukey makes a new Tukey window maker.
func NewTukey(alpha float64) (*Tukey, error) {
	if alpha < 0.0 || alpha > 1.0 {
		return nil, fmt.Errorf("alpha %f is not in range [0,1]", alpha)
	}

	return &Tukey{alpha: alpha}, nil
}

// Make Tukey window of a given size (number of samples).
func (w *Tukey) Make(size int) []float64 {
	coeff := make([]float64, size)

	sizeMinus1 := float64(size - 1)
	left := int(w.alpha * sizeMinus1 / 2.0)
	right := int(sizeMinus1 * (1.0 - (w.alpha / 2.0)))

	for n := 0; n < left; n++ {
		x := math.Pi * ((float64(2.0*n) / (w.alpha * sizeMinus1)) - 1.0)
		coeff[n] = 0.5 * (1.0 + math.Cos(x))
	}

	for n := left; n <= right; n++ {
		coeff[n] = 1.0
	}

	for n := right + 1; n < size; n++ {
		x := math.Pi * ((float64(2*n) / (w.alpha * sizeMinus1)) - (2.0 / w.alpha) + 1.0)
		coeff[n] = 0.5 * (1.0 + math.Cos(x))
	}

	return coeff
}
