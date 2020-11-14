package node_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	const val = 2.5

	n := node.NewConst(val)

	n.Initialize(100.0, 3)
	n.Run()

	assert.Equal(t, val, n.Output().Values[0])
	assert.Equal(t, val, n.Output().Values[1])
	assert.Equal(t, val, n.Output().Values[2])
}
