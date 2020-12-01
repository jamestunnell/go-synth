package window

// Rectangular is a Rectangular window maker.
// Produces a rectangular window (all ones) of a given size (number of samples).
// For more info, see https://en.wikipedia.org/wiki/Window_function#Rectangular_window.
type Rectangular struct {
}

// NewRectangular makes a new Rectangular window maker.
func NewRectangular() *Rectangular {
	return &Rectangular{}
}

// Make Rectangular window of a given size (number of samples).
func (w *Rectangular) Make(size int) []float64 {
	coeff := make([]float64, size)

	for n := 0; n < size; n++ {
		coeff[n] = 1.0
	}

	return coeff
}
