package window

import (
	"math"
)

// FlatTop is a FlatTop window maker.
// A flat top window is a partially negative-valued window that has a flat top in
// the frequency domain. They are designed to have a broader bandwidth and so have
// a poorer frequency resolution, leading to low amplitude measurement error suitable
// for use in in spectrum analyzers for the measurement of amplitudes of sinusoidal
// frequency components
// For more info, see https://en.wikipedia.org/wiki/Window_function//Flat_top_window
type FlatTop struct {
}

const (
	twoPi   = math.Pi * 2.0
	fourPi  = math.Pi * 4.0
	sixPi   = math.Pi * 6.0
	eightPi = math.Pi * 8.0
)

// NewFlatTop makes a new FlatTop window maker.
func NewFlatTop() *FlatTop {
	return &FlatTop{}
}

// Make FlatTop window of a given size (number of samples).
func (w *FlatTop) Make(size int) []float64 {
	const (
		a0 = 1.0
		a1 = 1.93
		a2 = 1.29
		a3 = 0.388
		a4 = 0.032
	)

	coeff := make([]float64, size)
	sizeMinus1 := float64(size - 1)
	max := math.Inf(-1)

	for n := 0; n < size; n++ {
		nFlt := float64(n)
		x := a0 -
			a1*math.Cos((twoPi*nFlt)/sizeMinus1) +
			a2*math.Cos((fourPi*nFlt)/sizeMinus1) -
			a3*math.Cos((sixPi*nFlt)/sizeMinus1) +
			a4*math.Cos((eightPi*nFlt)/sizeMinus1)

		coeff[n] = x

		if x > max {
			max = x
		}
	}

	oneOverMax := 1.0 / max

	//  normalize to max of 1.0
	for n := 0; n < size; n++ {
		coeff[n] *= oneOverMax
	}

	return coeff
}
