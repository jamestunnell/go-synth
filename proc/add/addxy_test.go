package add_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/generators/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/processors/add"
	"github.com/stretchr/testify/assert"
)

func TestAddXYHappyPath(t *testing.T) {
	in1 := array.OneShot([]float64{0.0, 0.1, 0.2})
	in2 := array.OneShot([]float64{-1.0, 0.5, -0.2})
	addXY := add.XY(in1, in2)

	node.Initialize(addXY, 100.0, 3)
	node.Run(addXY)

	assert.Equal(t, -1.0, addXY.Buffer().Values[0])
	assert.Equal(t, 0.6, addXY.Buffer().Values[1])
	assert.Equal(t, 0.0, addXY.Buffer().Values[2])
}
