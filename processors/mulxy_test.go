package processors_test

import (
	"testing"

	. "github.com/jamestunnell/go-synth/generators"
	"github.com/jamestunnell/go-synth/node"
	. "github.com/jamestunnell/go-synth/processors"
	"github.com/stretchr/testify/assert"
)

func TestMulXYHappyPath(t *testing.T) {
	in1 := &Array{Values: []float64{0.0, 0.1, 0.2}}
	in2 := &Array{Values: []float64{1.0, -1.0, 0.5}}

	mulXY := node.MakeNode(&MulXY{In1: in1, In2: in2}, 3)

	mulXY.Initialize(100.0)
	mulXY.Sample()

	assert.Equal(t, 0.0, mulXY.Out.Values[0])
	assert.Equal(t, -0.1, mulXY.Out.Values[1])
	assert.Equal(t, 0.1, mulXY.Out.Values[2])
}
