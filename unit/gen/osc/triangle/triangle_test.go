package triangle_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/osc/triangle"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestTriangle(t *testing.T) {
	f := node.NewConst(3.0)
	p := node.NewConst(0.0)
	n := triangle.NewNode(f, p)

	n.Initialize(15.0, 15)
	n.Run()

	outVals := n.Output().Values

	// First 5 samples should contain a complete cycle
	assert.InDelta(t, -1.0, outVals[0], 1e-5)
	assert.InDelta(t, -0.2, outVals[1], 1e-5)
	assert.InDelta(t, 0.6, outVals[2], 1e-5)
	assert.InDelta(t, 0.6, outVals[3], 1e-5)
	assert.InDelta(t, -0.2, outVals[4], 1e-5)

	// Then the first cycle should be repeated twice
	assert.Equal(t, outVals[:5], outVals[5:10])
	assert.Equal(t, outVals[:5], outVals[10:15])
}
