package fft_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/dsp/transform/fft"
	"github.com/jamestunnell/go-synth/util/complexslice"
)

func TestEnsurePowerOfTwoSizeInputSize1(t *testing.T) {
	input := []complex128{complex(2.5, 0.0)}
	actual, powerOfTwo := fft.EnsurePowerOfTwoSize(input)

	assert.Len(t, actual, 1)
	assert.Equal(t, 0, powerOfTwo)
	assert.Equal(t, input, actual)
}

func TestEnsurePowerOfTwoSizeInputSize2(t *testing.T) {
	input := []complex128{
		complex(2.5, 0.0),
		complex(-7.8, 0.0),
	}
	actual, powerOfTwo := fft.EnsurePowerOfTwoSize(input)

	assert.Len(t, actual, 2)
	assert.Equal(t, 1, powerOfTwo)
	assert.Equal(t, input, actual)
}

func TestEnsurePowerOfTwoSizeInputSize3(t *testing.T) {
	input := []complex128{
		complex(2.5, 0.0),
		complex(-7.8, 0.0),
		complex(29.8, 0.0),
	}
	actual, powerOfTwo := fft.EnsurePowerOfTwoSize(input)

	assert.Len(t, actual, 4)
	assert.Equal(t, 2, powerOfTwo)
	assert.Equal(t, input, actual[:3])
	assert.Equal(t, complexslice.Zeros(1), actual[3:4])
}
