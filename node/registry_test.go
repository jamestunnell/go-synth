package node_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/nodetest"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	r := node.WorkingRegistry()
	c := &nodetest.MulAdd{}
	path := node.CorePath(c)

	r.UnregisterCore(c)

	_, ok := r.MakeCore(path)

	assert.False(t, ok)

	r.RegisterCore(c)

	c2, ok := r.MakeCore(path)

	assert.True(t, ok)
	assert.NotNil(t, c2)
	assert.IsType(t, c, c2)
}
