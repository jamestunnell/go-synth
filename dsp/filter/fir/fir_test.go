package fir_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/dsp/filter/fir"
)

func TestNewFIR(t *testing.T) {
	f := fir.NewFIR([]float64{1.0, 0.0, 0.0})

	assert.Equal(t, 2, f.Order())
}
