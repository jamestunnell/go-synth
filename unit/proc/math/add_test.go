package math_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/jamestunnell/go-synth/unit/proc/math"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	in1 := array.NewOneshot([]float64{0.0, 0.1, 0.2})
	in2 := array.NewOneshot([]float64{-1.0, 0.5, -0.2})
	n := math.NewAdd(in1, in2)

	assert.NoError(t, n.Initialize(100.0, 3))

	n.Run()

	assert.Equal(t, -1.0, n.Output().Values[0])
	assert.Equal(t, 0.6, n.Output().Values[1])
	assert.Equal(t, 0.0, n.Output().Values[2])
}
