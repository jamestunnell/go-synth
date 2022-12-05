package synth_test

import (
	"testing"

	"github.com/jamestunnell/go-synth"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	r := synth.WorkingRegistry()
	c := &TestBlock{}
	path := synth.BlockPath(c)

	r.Unregister(path)

	_, ok := r.GetBlock(path)

	assert.False(t, ok)

	r.Register(c)

	c2, ok := r.GetBlock(path)

	assert.True(t, ok)
	assert.NotNil(t, c2)
	assert.IsType(t, c, c2)
}
