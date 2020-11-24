package node_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/nodetest"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	r := node.WorkingRegistry()
	c := &nodetest.TestCore{}
	path := node.CorePath(c)

	r.Unregister(path)

	_, ok := r.GetCore(path)

	assert.False(t, ok)

	r.Register(c)

	c2, ok := r.GetCore(path)

	assert.True(t, ok)
	assert.NotNil(t, c2)
	assert.IsType(t, c, c2)
}
