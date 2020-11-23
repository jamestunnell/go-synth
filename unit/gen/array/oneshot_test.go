package array_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/stretchr/testify/assert"
)

func TestOneshotNoValues(t *testing.T) {
	n := array.NewOneshot([]float64{})

	assert.Error(t, n.Initialize(100.0, 4))
}

func TestOneshotOneValueOneDeepBuffer(t *testing.T) {
	n := array.NewOneshot([]float64{2.5})

	assert.NoError(t, n.Initialize(100.0, 1))

	n.Run()

	assert.Equal(t, 2.5, n.Output().Values[0])

	n.Run()

	assert.Equal(t, 0.0, n.Output().Values[0])
}

func TestOneshotOneValueTwoDeepBuffer(t *testing.T) {
	n := array.NewOneshot([]float64{2.5})

	assert.NoError(t, n.Initialize(100.0, 2))

	n.Run()

	assert.Equal(t, 2.5, n.Output().Values[0])
	assert.Equal(t, 0.0, n.Output().Values[1])
}

func TestOneshotMultiValueOneDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	n := array.NewOneshot(vals)

	assert.NoError(t, n.Initialize(100.0, 1))

	for _, val := range vals {
		n.Run()

		assert.Equal(t, val, n.Output().Values[0])
	}

	n.Run()

	assert.Equal(t, 0.0, n.Output().Values[0])
}

func TestOneshotMultiValueMultiDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	n := array.NewOneshot(vals)

	assert.NoError(t, n.Initialize(100.0, len(vals)))

	n.Run()

	assert.Equal(t, vals, n.Output().Values)

	n.Run()

	assert.Equal(t, 0.0, n.Output().Values[0])
}
