package fft

import (
	"github.com/jamestunnell/go-synth/dsp/util/gain"
)

// FreqContent stores frequency magnitude and phase data.
type FreqContent struct {
	Frequencies []float64
	Magnitudes  []float64
	Phases      []float64
}

// MagnitudesDecibel returns the magnitude response in decibels.
func (fc *FreqContent) MagnitudesDecibel() ([]float64, error) {
	decibel := make([]float64, len(fc.Magnitudes))

	for i, mag := range fc.Magnitudes {
		dB, err := gain.LinearToDecibel(mag)
		if err != nil {
			return []float64{}, err
		}

		decibel[i] = dB
	}

	return decibel, nil
}
