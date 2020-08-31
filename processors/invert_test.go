package processors_test

import (
	"testing"

	. "github.com/jamestunnell/go-synth/generators"
	"github.com/jamestunnell/go-synth/node"
	. "github.com/jamestunnell/go-synth/processors"
	"github.com/stretchr/testify/assert"
)

func TestInvertHappyPath(t *testing.T) {
	in := &Array{Values: []float64{1.0, 0.5, -0.5}}
	inv := node.MakeNode(&Invert{In: in}, 3)

	inv.Initialize(100.0)
	inv.Sample()

	assert.Equal(t, 1.0, inv.Out.Values[0])
	assert.Equal(t, 2.0, inv.Out.Values[1])
	assert.Equal(t, -2.0, inv.Out.Values[2])
}
