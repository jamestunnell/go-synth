package window

import "math"

// Hamming is a Hamming window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
type Hamming struct {
}

// NewHamming makes a new Hamming window maker.
func NewHamming() *Hamming {
	return &Hamming{}
}

// Make Hamming window of a given size (number of samples).
func (w *Hamming) Make(size int) []float64 {
	const (
		alpha = 0.54
		beta  = 1.0 - alpha
	)

	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)

	for n := 0; n < size; n++ {
		theta := (twoPi * float64(n)) / sizeMinus1

		coeff[n] = alpha - (beta * math.Cos(theta))
	}

	return coeff
}
