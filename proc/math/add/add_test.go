package add_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/math/add"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	in1 := array.OneShot([]float64{0.0, 0.1, 0.2})
	in2 := array.OneShot([]float64{-1.0, 0.5, -0.2})
	a := add.New(in1, in2)

	node.Initialize(a, 100.0, 3)
	node.Run(a)

	assert.Equal(t, -1.0, a.Buffer().Values[0])
	assert.Equal(t, 0.6, a.Buffer().Values[1])
	assert.Equal(t, 0.0, a.Buffer().Values[2])
}
