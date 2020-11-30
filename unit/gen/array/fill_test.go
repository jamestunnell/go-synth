package array_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	x := []float64{0.0, 0.0, 0.0, 0.0, 0.0}

	array.Fill(x, 1.5)
	assert.Equal(t, []float64{1.5, 1.5, 1.5, 1.5, 1.5}, x)
}
