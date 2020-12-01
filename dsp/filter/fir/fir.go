package fir

import (
	"fmt"

	"github.com/jamestunnell/go-synth/dsp/transform/fft"
)

// FIR (Finite Impulse Response) filter
type FIR struct {
	kernel []float64
	order  int
}

// NewFIR makes a new FIR filter.
func NewFIR(kernel []float64) *FIR {
	return &FIR{
		kernel: kernel,
		order:  len(kernel) - 1,
	}
}

func (fir *FIR) Order() int {
	return fir.order
}

// Convolve the given input data with the filter kernel.
// Returns non-nil error if input size not be greater than the filter kernel size.
func (fir *FIR) Convolve(input []float64) ([]float64, error) {
	kernelSize := len(fir.kernel)
	n := len(input)

	if n <= kernelSize {
		err := fmt.Errorf("input size %d is not greater than kernel size %d", n, kernelSize)
		return []float64{}, err
	}

	output := make([]float64, n)

	//
	for i := 0; i < kernelSize; i++ {
		sum := 0.0

		// convolve the input with the filter kernel
		for j := 0; j < i; j++ {
			sum += (input[j] * fir.kernel[kernelSize-(1+i-j)])
		}

		output[i] = sum
	}

	for i := kernelSize; i < n; i++ {
		sum := 0.0

		// convolve the input with the filter kernel
		for j := 0; j < kernelSize; j++ {
			sum += (input[i-j] * fir.kernel[j])
		}

		output[i] = sum
	}

	return output, nil
}

// FreqResponse calculates the filter frequency response.
func (fir *FIR) FreqResponse(srate float64) *fft.FreqContent {
	return fft.AnalyzeFloats(srate, fir.kernel, fft.NoScaling)
}
