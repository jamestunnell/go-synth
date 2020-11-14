package saw_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/osc/saw"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestSawtooth(t *testing.T) {
	f := node.NewConst(3.0)
	p := node.NewConst(0.0)
	n := saw.NewNode(f, p)

	n.Initialize(15.0, 15)
	n.Run()

	outVals := n.Output().Values

	// First 5 samples should contain a complete cycle
	assert.Equal(t, 0.0, outVals[0])
	assert.Equal(t, 0.4, outVals[1])
	assert.Equal(t, 0.8, outVals[2])
	assert.Equal(t, -0.8, outVals[3])
	assert.Equal(t, -0.4, outVals[4])

	// Then the first cycle should be repeated twice
	assert.Equal(t, outVals[:5], outVals[5:10])
	assert.Equal(t, outVals[:5], outVals[10:15])
}
