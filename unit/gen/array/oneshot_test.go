package array_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/stretchr/testify/assert"
)

func TestOneshotNoValues(t *testing.T) {
	o := array.NewOneshot()

	assert.Error(t, o.Initialize(100.0, 4))
}

func TestOneshotOneValueOneDeepBuffer(t *testing.T) {
	o := array.NewOneshot()

	assert.True(t, o.Values.SetValue([]float64{2.5}))

	assert.NoError(t, o.Initialize(100.0, 1))

	o.Run()

	outVals := o.Out.Buffer().([]float64)

	assert.Equal(t, 2.5, outVals[0])

	o.Run()

	assert.Equal(t, 0.0, outVals[0])
}

func TestOneshotOneValueTwoDeepBuffer(t *testing.T) {
	o := array.NewOneshot()

	assert.True(t, o.Values.SetValue([]float64{2.5}))

	assert.NoError(t, o.Initialize(100.0, 2))

	o.Run()

	outVals := o.Out.Buffer().([]float64)

	assert.Equal(t, 2.5, outVals[0])
	assert.Equal(t, 0.0, outVals[1])
}

func TestOneshotMultiValueOneDeepBuffer(t *testing.T) {
	o := array.NewOneshot()
	vals := []float64{0.3, 2.2, -4.5, 66.88}

	assert.True(t, o.Values.SetValue(vals))

	assert.NoError(t, o.Initialize(100.0, 1))

	outVals := o.Out.Buffer().([]float64)

	for _, val := range vals {
		o.Run()

		assert.Equal(t, val, outVals[0])
	}

	o.Run()

	assert.Equal(t, 0.0, outVals[0])
}

func TestOneshotMultiValueMultiDeepBuffer(t *testing.T) {
	o := array.NewOneshot()
	vals := []float64{0.3, 2.2, -4.5, 66.88}

	assert.True(t, o.Values.SetValue(vals))

	assert.NoError(t, o.Initialize(100.0, len(vals)))

	outVals := o.Out.Buffer().([]float64)

	o.Run()

	assert.Equal(t, vals, outVals)

	o.Run()

	assert.Equal(t, 0.0, outVals[0])
}
