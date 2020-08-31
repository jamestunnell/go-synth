package processors_test

import (
	"testing"

	. "github.com/jamestunnell/go-synth/generators"
	"github.com/jamestunnell/go-synth/node"
	. "github.com/jamestunnell/go-synth/processors"
	"github.com/stretchr/testify/assert"
)

func TestAddKHappyPath(t *testing.T) {
	in := &Array{Values: []float64{0.0, 0.1, 0.2}}
	addK := node.MakeNode(&AddK{In: in, K: 1.0}, 3)

	addK.Initialize(100.0)
	addK.Sample()

	assert.Equal(t, 1.0, addK.Out.Values[0])
	assert.Equal(t, 1.1, addK.Out.Values[1])
	assert.Equal(t, 1.2, addK.Out.Values[2])
}
