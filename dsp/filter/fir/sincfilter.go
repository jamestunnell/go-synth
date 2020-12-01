package fir

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-synth/dsp/transform/fft"
	"github.com/jamestunnell/go-synth/dsp/window"
)

// SincFilter is windowed FIR filter that implements lowpass and highpass.
// A bandpass and bandstop filter would be implemented using two of these.
// Theoretical source: http://www.labbookpages.co.uk/audio/firWindowing.html
type SincFilter struct {
	srate  float64
	order  int
	cutoff float64
	// window []float64

	lowpassFIR  *FIR
	highpassFIR *FIR
}

const twoPi = math.Pi * 2.0

// NewSincFilter makes a new FIR filter that can be used for highpass and lowpass filtering.
func NewSincFilter(srate, cutoff float64, order int, w window.WindowMaker) (*SincFilter, error) {
	if srate <= 0.0 {
		err := fmt.Errorf("sample rate %f is not positive", srate)
		return nil, err
	}

	if cutoff > (srate / 2.0) {
		err := fmt.Errorf("cutoff %f is greater than half sample rate %f", cutoff, srate)
		return nil, err
	}

	if order == 0 {
		err := fmt.Errorf("order %d is not positive", order)
		return nil, err
	}

	if (order % 2) != 0 {
		err := fmt.Errorf("order %d is not even", order)
		return nil, err
	}

	size := order + 1
	transitionFreq := cutoff / srate
	b := twoPi * transitionFreq

	// make FIR filter kernels for lowpass and highpass
	lpKernel := make([]float64, size)
	hpKernel := make([]float64, size)
	win := w.Make(size)

	for n := 0; n < (order / 2); n++ {
		c := float64(n - (order / 2))
		y := win[n] * math.Sin(b*c) / (math.Pi * c)

		lpKernel[n] = y
		lpKernel[size-1-n] = y

		hpKernel[n] = -y
		hpKernel[size-1-n] = hpKernel[n]
	}

	lpKernel[order/2] = 2 * transitionFreq * win[order/2]
	hpKernel[order/2] = (1 - 2*transitionFreq) * win[order/2]

	f := &SincFilter{
		srate:       srate,
		order:       order,
		cutoff:      cutoff,
		lowpassFIR:  NewFIR(lpKernel),
		highpassFIR: NewFIR(hpKernel),
	}

	return f, nil
}

// SampleRate returns the sample rate that the filter was designed with.
func (f *SincFilter) SampleRate() float64 {
	return f.srate
}

// Lowpass processes the input with the lowpass FIR.
func (f *SincFilter) Lowpass(input []float64) ([]float64, error) {
	return f.lowpassFIR.Convolve(input)
}

// Highpass processes the input with the highpass FIR.
func (f *SincFilter) Highpass(input []float64) ([]float64, error) {
	return f.lowpassFIR.Convolve(input)
}

// LowpassResponse returns the frequency response of the lowpass FIR.
func (f *SincFilter) LowpassResponse() *fft.FreqContent {
	return f.lowpassFIR.FreqResponse(f.srate)
}

// HighpassResponse returns the frequency response of the highpass FIR.
func (f *SincFilter) HighpassResponse() *fft.FreqContent {
	return f.highpassFIR.FreqResponse(f.srate)
}
