package array_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/stretchr/testify/assert"
)

func TestRepeatNoValues(t *testing.T) {
	r := array.NewRepeat()

	assert.Error(t, r.Initialize(100.0, 4))
}

func TestRepeatMultiValueOneDeepBuffer(t *testing.T) {
	r := array.NewRepeat()
	vals := []float64{2.5, 3.3}

	assert.True(t, r.Values.SetValue(vals))

	assert.NoError(t, r.Initialize(100.0, 1))

	r.Run()

	assert.Equal(t, vals[0], r.Out.BufferValues[0])

	r.Run()

	assert.Equal(t, vals[1], r.Out.BufferValues[0])

	r.Run()

	assert.Equal(t, vals[0], r.Out.BufferValues[0])

	r.Run()

	assert.Equal(t, vals[1], r.Out.BufferValues[0])
}

func TestRepeatOneValueTwoDeepBuffer(t *testing.T) {
	r := array.NewRepeat()

	assert.True(t, r.Values.SetValue([]float64{2.5}))

	assert.NoError(t, r.Initialize(100.0, 2))

	r.Run()

	assert.Equal(t, 2.5, r.Out.BufferValues[0])
	assert.Equal(t, 2.5, r.Out.BufferValues[1])
}

func TestRepeatMultiValueOddSizeBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2}
	r := array.NewRepeat()

	assert.True(t, r.Values.SetValue(vals))

	assert.NoError(t, r.Initialize(100.0, 3))

	r.Run()

	assert.Equal(t, vals[0], r.Out.BufferValues[0])
	assert.Equal(t, vals[1], r.Out.BufferValues[1])
	assert.Equal(t, vals[0], r.Out.BufferValues[2])

	r.Run()

	assert.Equal(t, vals[1], r.Out.BufferValues[0])
	assert.Equal(t, vals[0], r.Out.BufferValues[1])
	assert.Equal(t, vals[1], r.Out.BufferValues[2])

	r.Run()

	assert.Equal(t, vals[0], r.Out.BufferValues[0])
	assert.Equal(t, vals[1], r.Out.BufferValues[1])
	assert.Equal(t, vals[0], r.Out.BufferValues[2])
}
