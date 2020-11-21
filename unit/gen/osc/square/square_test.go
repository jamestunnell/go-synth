package square_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/osc/square"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	f := node.NewConst(3.0)
	p := node.NewConst(0.0)
	n := square.NewNode(f, p)

	n.Initialize(15.0, 15)
	n.Run()

	outVals := n.Output().Values

	// First 5 samples should contain a complete cycle
	assert.Equal(t, 1.0, outVals[0])
	assert.InDelta(t, 1.0, outVals[1], 1e-3)
	assert.InDelta(t, 1.0, outVals[2], 1e-3)
	assert.InDelta(t, -1.0, outVals[3], 1e-3)
	assert.InDelta(t, -1.0, outVals[4], 1e-3)

	// Then the first cycle should be repeated twice
	assert.Equal(t, outVals[:5], outVals[5:10])
	assert.Equal(t, outVals[:5], outVals[10:15])
}
