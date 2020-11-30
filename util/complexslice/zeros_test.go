package complexslice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/util/complexslice"
)

func TestZeros(t *testing.T) {
	zero := complex(0.0, 0.0)
	actual := complexslice.Zeros(5)
	expected := []complex128{zero, zero, zero, zero, zero}

	assert.Equal(t, expected, actual)
}
