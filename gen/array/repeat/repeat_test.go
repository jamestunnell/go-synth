package repeat_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array/repeat"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestRepeatNoValues(t *testing.T) {
	defer func() { recover() }()

	repeat.New([]float64{})

	t.Errorf("did not panic")
}

func TestRepeatOneValueOneDeepBuffer(t *testing.T) {
	n := repeat.New([]float64{2.5})
	out := node.NewBuffer(1)

	n.Run(out)

	assert.Equal(t, 2.5, out.Values[0])

	n.Run(out)

	assert.Equal(t, 2.5, out.Values[0])
}

func TestRepeatOneValueTwoDeepBuffer(t *testing.T) {
	n := repeat.New([]float64{2.5})
	out := node.NewBuffer(2)

	n.Run(out)

	assert.Equal(t, 2.5, out.Values[0])
	assert.Equal(t, 2.5, out.Values[1])
}

func TestRepeatMultiValueOneDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	n := repeat.New(vals)
	out := node.NewBuffer(1)

	for i := 0; i < 3; i++ {
		for _, val := range vals {
			n.Run(out)

			assert.Equal(t, val, out.Values[0])
		}
	}
}

func TestRepeatMultiValueMultiDeepBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2, -4.5, 66.88}
	n := repeat.New(vals)
	out := node.NewBuffer(len(vals))

	n.Run(out)
	assert.Equal(t, vals, out.Values)

	n.Run(out)
	assert.Equal(t, vals, out.Values)
}
