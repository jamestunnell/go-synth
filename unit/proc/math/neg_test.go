package math_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/jamestunnell/go-synth/unit/proc/math"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNegateHappyPath(t *testing.T) {
	in := array.NewOneshot()
	inVals := []float64{1.0, 0.5, -0.5}

	require.NoError(t, in.Values.SetValue(inVals))

	a := math.NewNeg()

	a.In.Connect(in.Out)

	require.NoError(t, in.Initialize(100.0, 3))
	require.NoError(t, a.Initialize(100.0, 3))

	in.Run()
	a.Run()

	assert.Equal(t, -1.0, a.Out.BufferValues[0])
	assert.Equal(t, -0.5, a.Out.BufferValues[1])
	assert.Equal(t, 0.5, a.Out.BufferValues[2])
}
