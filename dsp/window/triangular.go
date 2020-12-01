package window

import "math"

// Triangular is a Triangular window maker.
// Endpoints are near zero. Midpoint is one. There is a linear slope between endpoints and midpoint.
// For more info, see https://en.wikipedia.org/wiki/Window_function//Triangular_window
type Triangular struct {
}

// NewTriangular makes a new Triangular window maker.
func NewTriangular() *Triangular {
	return &Triangular{}
}

// Make Triangular window of a given size (number of samples).
func (w *Triangular) Make(size int) []float64 {
	coeff := make([]float64, size)
	sizePlus1 := float64(size + 1)
	sizeMinus1Over2 := float64(size-1) / 2.0
	sizePlus1Over2 := sizePlus1 / 2.0

	for n := 0; n < size; n++ {
		x := sizePlus1Over2 - math.Abs(float64(n)-sizeMinus1Over2)

		coeff[n] = (2.0 / sizePlus1) * x
	}

	return coeff
}
