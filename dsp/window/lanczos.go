package window

import "math"

// Lanczos is a Lanczos window maker.
// The Lanczos window is used in Lanczos resampling.
// For more info, see https://en.wikipedia.org/wiki/Window_function#Lanczos_window.
type Lanczos struct {
}

// NewLanczos makes a new Lanczos window maker.
func NewLanczos() *Lanczos {
	return &Lanczos{}
}

// Make Lanczos window of a given size (number of samples).
func (w *Lanczos) Make(size int) []float64 {
	sinc := func(x float64) float64 {
		return math.Sin(math.Pi*x) / (math.Pi * x)
	}
	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)

	for n := 0; n < size; n++ {
		x := (2.0 * float64(n) / sizeMinus1) - 1.0
		coeff[n] = sinc(x)
	}

	return coeff
}
