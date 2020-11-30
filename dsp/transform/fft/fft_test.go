package fft_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/dsp/transform/fft"
	"github.com/jamestunnell/go-synth/util/complexslice"
	"github.com/stretchr/testify/assert"
)

// TestForwardWithImpulse sends an impulse through the forward FFT.
// See http://www.sccon.ca/sccon/fft/fft3.htm
func TestForwardWithImpulse(t *testing.T) {
	impulse := append([]complex128{complex(1.0, 0.0)}, complexslice.Zeros(7)...)
	expected := []complex128{
		complex(0.125, 0.0),
		complex(0.125, 0.0),
		complex(0.125, 0.0),
		complex(0.125, 0.0),
		complex(0.125, 0.0),
		complex(0.125, 0.0),
		complex(0.125, 0.0),
		complex(0.125, 0.0),
	}
	actual, err := fft.Forward(impulse)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

// TestForwardWithShiftedImpulse sends a shifted impulse through the forward FFT.
// See http://www.sccon.ca/sccon/fft/fft3.htm
func TestForwardWithShiftedImpulse(t *testing.T) {
	impulse := append(
		[]complex128{complex(0.0, 0.0), complex(1.0, 0.0)},
		complexslice.Zeros(6)...,
	)
	expected := []complex128{
		complex(0.125, 0.0),
		complex(0.088388, -0.088388),
		complex(0.0, -0.125),
		complex(-0.088388, -0.088388),
		complex(-0.125, 0.0),
		complex(-0.088388, 0.088388),
		complex(0.0, 0.125),
		complex(0.088388, 0.088388),
	}
	actual, err := fft.Forward(impulse)

	assert.NoError(t, err)
	verifyEqualWithin(t, expected, actual, 1e-4)
}

// TestForwardInverse
func TestForwardInverse(t *testing.T) {
	impulse := append(
		[]complex128{complex(1.0, 0.0)},
		complexslice.Zeros(7)...,
	)
	actual, err := fft.Forward(impulse)

	assert.NoError(t, err)

	actual, err = fft.Inverse(actual)

	assert.NoError(t, err)
	verifyEqualWithin(t, impulse, actual, 1e-10)
}

func verifyEqualWithin(t *testing.T, expected, actual []complex128, delta float64) {
	if !assert.Len(t, actual, len(expected)) {
		return
	}

	for i := 0; i < len(actual); i++ {
		assert.InDelta(t, real(expected[i]), real(actual[i]), delta)
		assert.InDelta(t, imag(expected[i]), imag(actual[i]), delta)
	}
}
