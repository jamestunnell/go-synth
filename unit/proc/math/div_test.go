package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/jamestunnell/go-synth/unit/proc/math"
)

func TestDivHappyPath(t *testing.T) {
	in1Vals := []float64{0.0, 0.1, 0.2, 1.0}
	in2Vals := []float64{1.0, -1.0, 0.5, 2.0}

	in1 := array.NewOneshot()
	in2 := array.NewOneshot()
	blk := math.NewDiv()

	assert.NoError(t, in1.Values.SetValue(in1Vals))
	assert.NoError(t, in2.Values.SetValue(in2Vals))
	assert.NoError(t, blk.In1.Connect(in1.Out))
	assert.NoError(t, blk.In2.Connect(in2.Out))

	require.NoError(t, in1.Initialize(100.0, 4))
	require.NoError(t, in2.Initialize(100.0, 4))
	assert.NoError(t, blk.Initialize(100.0, 4))

	in1.Run()
	in2.Run()
	blk.Run()

	assert.Equal(t, 0.0, blk.Out.BufferValues[0])
	assert.Equal(t, -0.1, blk.Out.BufferValues[1])
	assert.Equal(t, 0.4, blk.Out.BufferValues[2])
	assert.Equal(t, 0.5, blk.Out.BufferValues[3])
}
