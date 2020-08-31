package processors_test

import (
	"testing"

	. "github.com/jamestunnell/go-synth/generators"
	"github.com/jamestunnell/go-synth/node"
	. "github.com/jamestunnell/go-synth/processors"
	"github.com/stretchr/testify/assert"
)

func TestmulKHappyPath(t *testing.T) {
	in := &Array{Values: []float64{0.0, 0.2, -0.3}}
	mulK := node.MakeNode(&MulK{In: in, K: 2.0}, 3)

	mulK.Initialize(100.0)
	mulK.Sample()

	assert.Equal(t, 0.0, mulK.Out.Values[0])
	assert.Equal(t, 0.4, mulK.Out.Values[1])
	assert.Equal(t, -0.6, mulK.Out.Values[2])
}
