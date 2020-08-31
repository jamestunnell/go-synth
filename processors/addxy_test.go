package processors_test

import (
	"testing"

	. "github.com/jamestunnell/go-synth/generators"
	"github.com/jamestunnell/go-synth/node"
	. "github.com/jamestunnell/go-synth/processors"
	"github.com/stretchr/testify/assert"
)

func TestAddXYHappyPath(t *testing.T) {
	in1 := &Array{Values: []float64{0.0, 0.1, 0.2}}
	in2 := &Array{Values: []float64{-1.0, 0.5, -0.2}}

	addXY := node.MakeNode(&AddXY{In1: in1, In2: in2}, 3)

	addXY.Initialize(100.0)
	addXY.Sample()

	assert.Equal(t, -1.0, addXY.Out.Values[0])
	assert.Equal(t, 0.6, addXY.Out.Values[1])
	assert.Equal(t, 0.0, addXY.Out.Values[2])
}
