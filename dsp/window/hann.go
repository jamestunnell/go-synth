package window

import "math"

// Hann is a Hann window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
type Hann struct {
}

// NewHann makes a new Hann window maker.
func NewHann() *Hann {
	return &Hann{}
}

// Make Hann window of a given size (number of samples).
func (w *Hann) Make(size int) []float64 {
	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)

	for n := 0; n < size; n++ {
		theta := (twoPi * float64(n)) / sizeMinus1

		coeff[n] = 0.5 * (1.0 - math.Cos(theta))
	}

	return coeff
}
