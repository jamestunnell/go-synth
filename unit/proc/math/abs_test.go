package math_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/jamestunnell/go-synth/unit/proc/math"
	"github.com/stretchr/testify/assert"
)

func TestAbsHappyPath(t *testing.T) {
	in := array.NewOneshot([]float64{1.0, 0.5, -0.5})
	n := math.NewAbs(in)

	assert.NoError(t, n.Initialize(100.0, 3))

	n.Run()

	assert.Equal(t, 1.0, n.Output().Values[0])
	assert.Equal(t, 0.5, n.Output().Values[1])
	assert.Equal(t, 0.5, n.Output().Values[2])
}
