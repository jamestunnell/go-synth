package fft

import (
	"math"

	"github.com/jamestunnell/go-synth/util/complexslice"
)

func EnsurePowerOfTwoSize(vals []complex128) ([]complex128, int) {
	size := len(vals)
	powerOfTwo := math.Log2(float64(size))
	powerOfTwoFloor := math.Floor(powerOfTwo)

	// ensure input size is an even exact of two
	if powerOfTwoFloor != powerOfTwo {
		nextPowerOfTwo := powerOfTwoFloor + 1.0
		newSize := int(math.Pow(2, nextPowerOfTwo))

		vals = append(vals, complexslice.Zeros(newSize-size)...)
		size = newSize
		powerOfTwo = nextPowerOfTwo
	}

	return vals, int(powerOfTwo)
}
