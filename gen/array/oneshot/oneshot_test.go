package oneshot_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestOneshotNoValues(t *testing.T) {
	defer func() { recover() }()

	oneshot.New([]float64{})

	t.Errorf("did not panic")
}

func TestOneshotOneValueOneDeepBuffer(t *testing.T) {
	n := oneshot.New([]float64{2.5})
	out := node.NewBuffer(1)

	n.Run(out)

	assert.Equal(t, 2.5, out.Values[0])

	n.Run(out)

	assert.Equal(t, 0.0, out.Values[0])
}

func TestOneshotOneValueTwoDeepBuffer(t *testing.T) {
	n := oneshot.New([]float64{2.5})
	out := node.NewBuffer(2)

	n.Run(out)

	assert.Equal(t, 2.5, out.Values[0])
	assert.Equal(t, 0.0, out.Values[1])
}

func TestOneshotMultiValueOneDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	n := oneshot.New(vals)
	out := node.NewBuffer(1)

	for _, val := range vals {
		n.Run(out)

		assert.Equal(t, val, out.Values[0])
	}

	n.Run(out)

	assert.Equal(t, 0.0, out.Values[0])
}

func TestOneshotMultiValueMultiDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	n := oneshot.New(vals)
	out := node.NewBuffer(len(vals))

	n.Run(out)
	assert.Equal(t, vals, out.Values)

	n.Run(out)

	assert.Equal(t, 0.0, out.Values[0])
}
