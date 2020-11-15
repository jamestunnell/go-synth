package repeat_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array/repeat"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestRepeatNoValues(t *testing.T) {
	r := repeat.New([]float64{})
	out := node.NewBuffer(4)

	r.Run(out)

	assert.Equal(t, []float64{0.0, 0.0, 0.0, 0.0}, out.Values)
}

func TestRepeatMultiValueOneDeepBuffer(t *testing.T) {
	vals := []float64{2.5, 3.3}
	r := repeat.New(vals)
	out := node.NewBuffer(1)

	r.Run(out)

	assert.Equal(t, vals[0], out.Values[0])

	r.Run(out)

	assert.Equal(t, vals[1], out.Values[0])

	r.Run(out)

	assert.Equal(t, vals[0], out.Values[0])

	r.Run(out)

	assert.Equal(t, vals[1], out.Values[0])
}

func TestRepeatOneValueTwoDeepBuffer(t *testing.T) {
	r := repeat.New([]float64{2.5})
	out := node.NewBuffer(2)

	r.Run(out)

	assert.Equal(t, 2.5, out.Values[0])
	assert.Equal(t, 2.5, out.Values[1])
}

func TestRepeatMultiValueOddSizeBuffer(t *testing.T) {
	vals := []float64{0.3, 2.2}
	r := repeat.New(vals)
	out := node.NewBuffer(3)

	r.Run(out)
	assert.Equal(t, vals[0], out.Values[0])
	assert.Equal(t, vals[1], out.Values[1])
	assert.Equal(t, vals[0], out.Values[2])

	r.Run(out)
	assert.Equal(t, vals[1], out.Values[0])
	assert.Equal(t, vals[0], out.Values[1])
	assert.Equal(t, vals[1], out.Values[2])

	r.Run(out)
	assert.Equal(t, vals[0], out.Values[0])
	assert.Equal(t, vals[1], out.Values[1])
	assert.Equal(t, vals[0], out.Values[2])
}
