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

	assert.NoError(t, o.Values.SetValue([]float64{2.5}))

	assert.NoError(t, o.Initialize(100.0, 1))

	o.Run()

	assert.Equal(t, 2.5, o.Out.Buffer[0])

	o.Run()

	assert.Equal(t, 0.0, o.Out.Buffer[0])
}

func TestOneshotOneValueTwoDeepBuffer(t *testing.T) {
	o := array.NewOneshot()

	assert.NoError(t, o.Values.SetValue([]float64{2.5}))

	assert.NoError(t, o.Initialize(100.0, 2))

	o.Run()

	assert.Equal(t, 2.5, o.Out.Buffer[0])
	assert.Equal(t, 0.0, o.Out.Buffer[1])
}

func TestOneshotMultiValueOneDeepBuffer(t *testing.T) {
	o := array.NewOneshot()
	vals := []float64{0.3, 2.2, -4.5, 66.88}

	assert.NoError(t, o.Values.SetValue(vals))

	assert.NoError(t, o.Initialize(100.0, 1))

	for _, val := range vals {
		o.Run()

		assert.Equal(t, val, o.Out.Buffer[0])
	}

	o.Run()

	assert.Equal(t, 0.0, o.Out.Buffer[0])
}

func TestOneshotMultiValueMultiDeepBuffer(t *testing.T) {
	o := array.NewOneshot()
	vals := []float64{0.3, 2.2, -4.5, 66.88}

	assert.NoError(t, o.Values.SetValue(vals))

	assert.NoError(t, o.Initialize(100.0, len(vals)))

	o.Run()

	assert.Equal(t, vals, o.Out.Buffer)

	o.Run()

	assert.Equal(t, 0.0, o.Out.Buffer[0])
}
