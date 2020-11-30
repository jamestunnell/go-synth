package fft

import (
	"fmt"
	"math"
)

const twoPi = math.Pi * 2.0

// Forward Radix-2 FFT transform using decimation-in-time.
// Returns non-nil error if input size is not an exact power of two.
// EnsurePowerOfTwoSize can be used to make power of two size by padding with zeros.
// Ported from unlicensed MATLAB code which was posted to the MathWorks file
// exchange by Dinesh Dileep Gaurav.
// See http://www.mathworks.com/matlabcentral/fileexchange/17778.
func Forward(vals []complex128) ([]complex128, error) {
	x, err := fftCommon(vals)
	if err != nil {
		return []complex128{}, err
	}

	// scale the output values by the input size
	size := len(x)
	scale := complex(1.0/float64(size), 0.0)

	for i := 0; i < size; i++ {
		x[i] *= scale
	}

	return x, nil
}

// Inverse Radix-2 FFT transform.
// Returns non-nil error if input size is not an exact power of two.
// Ported from unlicensed MATLAB code which was posted to the MathWorks file
// exchange by Dinesh Dileep Gaurav.
// See http://www.mathworks.com/matlabcentral/fileexchange/17778.
func Inverse(vals []complex128) ([]complex128, error) {
	return fftCommon(vals)
}

func fftCommon(vals []complex128) ([]complex128, error) {
	size := len(vals)
	powerOfTwo := math.Log2(float64(size))

	if math.Floor(powerOfTwo) != powerOfTwo {
		err := fmt.Errorf("input size %d is not a power of 2", size)
		return []complex128{}, err
	}

	x := bitReversedOrder(vals, int(powerOfTwo))

	phase := make([]complex128, size/2)

	for i := 0; i < len(phase); i++ {
		theta := twoPi * float64(i) / float64(size)
		phase[i] = complex(math.Cos(theta), -math.Sin(theta))
	}

	for a := 1; a <= int(powerOfTwo); a++ {
		l := (1 << a) // 2^a
		phaseLevel := []complex128{}

		for i := 0; i < size/2; i += (size / l) {
			phaseLevel = append(phaseLevel, phase[i])
		}

		for k := 0; k <= (size - l); k += l {
			for n := 0; n < (l / 2); n++ {
				idx1 := n + k
				idx2 := n + k + (l / 2)

				first := x[idx1]
				second := x[idx2] * phaseLevel[n]
				x[idx1] = first + second
				x[idx2] = first - second
			}
		}
	}

	return x, nil
}

// bitReversedOrder reorders the input values using bit reversed indices.
func bitReversedOrder(vals []complex128, nBits int) []complex128 {
	n := len(vals)
	newVals := make([]complex128, n)

	for i := uint64(0); i < uint64(n); i++ {
		newIdx, err := BitReverse(i, nBits)

		// We don't expect this to ever fail
		if err != nil {
			err = fmt.Errorf("failed to bit-reverse index %d: %v", i, err)
			panic(err)
		}

		newVals[i] = vals[newIdx]
	}

	return newVals
}
