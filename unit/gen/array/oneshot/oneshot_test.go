package oneshot_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestOneshotNoValues(t *testing.T) {
	o := oneshot.New([]float64{})
	out := node.NewBuffer(4)

	o.Run(out)

	assert.Equal(t, []float64{0.0, 0.0, 0.0, 0.0}, out.Values)
}

func TestOneshotOneValueOneDeepBuffer(t *testing.T) {
	o := oneshot.New([]float64{2.5})
	out := node.NewBuffer(1)

	o.Run(out)

	assert.Equal(t, 2.5, out.Values[0])

	o.Run(out)

	assert.Equal(t, 0.0, out.Values[0])
}

func TestOneshotOneValueTwoDeepBuffer(t *testing.T) {
	o := oneshot.New([]float64{2.5})
	out := node.NewBuffer(2)

	o.Run(out)

	assert.Equal(t, 2.5, out.Values[0])
	assert.Equal(t, 0.0, out.Values[1])
}

func TestOneshotMultiValueOneDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	o := oneshot.New(vals)
	out := node.NewBuffer(1)

	for _, val := range vals {
		o.Run(out)

		assert.Equal(t, val, out.Values[0])
	}

	o.Run(out)

	assert.Equal(t, 0.0, out.Values[0])
}

func TestOneshotMultiValueMultiDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	o := oneshot.New(vals)
	out := node.NewBuffer(len(vals))

	o.Run(out)
	assert.Equal(t, vals, out.Values)

	o.Run(out)

	assert.Equal(t, 0.0, out.Values[0])
}
