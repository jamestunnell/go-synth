package window

import (
	"fmt"
	"math"
)

// Gaussian is a Gaussian window maker.
// For more info, see
// https://en.wikipedia.org/wiki/Window_function#Gaussian_windows
type Gaussian struct {
	sigma float64
}

// NewGaussian makes a new Gaussian window maker.
// Sigma must be <= 0.5.
func NewGaussian(sigma float64) (*Gaussian, error) {
	if sigma < 0.0 || sigma > 0.5 {
		return nil, fmt.Errorf("sigma %f is not in range [0,1]", sigma)
	}

	return &Gaussian{sigma: sigma}, nil
}

// Make Gaussian window of a given size (number of samples).
func (w *Gaussian) Make(size int) []float64 {
	coeff := make([]float64, size)
	sizeOver2 := float64(size / 2)

	for n := 0; n < size; n++ {
		a := (float64(n) - sizeOver2) / (w.sigma * sizeOver2)

		coeff[n] = math.Exp(-0.5 * a * a)
	}

	return coeff
}
