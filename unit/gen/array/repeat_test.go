package array_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/stretchr/testify/assert"
)

func TestRepeatNoValues(t *testing.T) {
	n := array.NewRepeat([]float64{})

	assert.Error(t, n.Initialize(100.0, 4))
}

func TestRepeatMultiValueOneDeepBuffer(t *testing.T) {
	vals := []float64{2.5, 3.3}
	n := array.NewRepeat(vals)

	assert.NoError(t, n.Initialize(100.0, 1))

	n.Run()

	assert.Equal(t, vals[0], n.Output().Values[0])

	n.Run()

	assert.Equal(t, vals[1], n.Output().Values[0])

	n.Run()

	assert.Equal(t, vals[0], n.Output().Values[0])

	n.Run()

	assert.Equal(t, vals[1], n.Output().Values[0])
}

func TestRepeatOneValueTwoDeepBuffer(t *testing.T) {
	n := array.NewRepeat([]float64{2.5})

	assert.NoError(t, n.Initialize(100.0, 2))

	n.Run()

	assert.Equal(t, 2.5, n.Output().Values[0])
	assert.Equal(t, 2.5, n.Output().Values[1])
}

func TestRepeatMultiValueOddSizeBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2}
	n := array.NewRepeat(vals)

	assert.NoError(t, n.Initialize(100.0, 3))

	n.Run()

	assert.Equal(t, vals[0], n.Output().Values[0])
	assert.Equal(t, vals[1], n.Output().Values[1])
	assert.Equal(t, vals[0], n.Output().Values[2])

	n.Run()

	assert.Equal(t, vals[1], n.Output().Values[0])
	assert.Equal(t, vals[0], n.Output().Values[1])
	assert.Equal(t, vals[1], n.Output().Values[2])

	n.Run()

	assert.Equal(t, vals[0], n.Output().Values[0])
	assert.Equal(t, vals[1], n.Output().Values[1])
	assert.Equal(t, vals[0], n.Output().Values[2])
}
