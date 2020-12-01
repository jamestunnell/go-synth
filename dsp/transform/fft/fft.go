package fft

import (
	"fmt"
	"math"
	"math/cmplx"

	"github.com/jamestunnell/go-synth/util/complexslice"
)

type Scaling int

const (
	NoScaling Scaling = iota
	ScaleByOneOverN
	ScaleByOneOverSqrtN

	twoPi = math.Pi * 2.0
)

// FFT is a radix-2 FFT transform using decimation-in-time.
// Can be used for both forward (anaysis) and inverse (synthesis) transform
// by selecting appropriate scaling.
// Returns non-nil error if input size is not an exact power of two.
// EnsurePowerOfTwoSize can be used before forward FFT to make power of two
//  size by padding with zeros.
// Ported from unlicensed MATLAB code which was posted to the MathWorks file
// exchange by Dinesh Dileep Gaurav.
// See http://www.mathworks.com/matlabcentral/fileexchange/17778.
func FFT(vals []complex128, scaling Scaling) ([]complex128, error) {
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

	scale := complex(1.0, 0.0)
	switch scaling {
	case NoScaling:
		return x, nil
	case ScaleByOneOverN:
		scale = complex(1.0/float64(size), 0.0)
	case ScaleByOneOverSqrtN:
		scale = complex(1.0/math.Sqrt(float64(size)), 0.0)
	}

	for i := 0; i < size; i++ {
		x[i] *= scale
	}

	return x, nil
}

// AnalyzeFloats perfoms FFT on the given float values and returns frequency content.
// Before running the FFT, the float values will be padded with zeros to make radix-2 length.
// Only the first half of the FFT results (positive frequencies) will be included in the
// frequency content.
func AnalyzeFloats(srate float64, floatVals []float64, scaling Scaling) *FreqContent {
	input := complexslice.FromFloats(floatVals)
	input, _ = EnsurePowerOfTwoSize(input)

	output, _ := FFT(input, scaling)
	size := len(output)
	sizeHalf := size / 2

	// calculate magnitude response of first half (second half is a mirror image)
	mags := make([]float64, sizeHalf)
	phases := make([]float64, sizeHalf)
	freqs := make([]float64, sizeHalf)
	binScale := srate / float64(size)

	for i := 0; i < sizeHalf; i++ {
		mags[i], phases[i] = cmplx.Polar(output[i])
		freqs[i] = float64(i) * binScale
	}

	return &FreqContent{
		Frequencies: freqs,
		Magnitudes:  mags,
		Phases:      phases,
	}
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
