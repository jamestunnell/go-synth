package synth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth"
)

func TestRegistry(t *testing.T) {
	r := synth.WorkingRegistry()
	c := &TestBlock{}
	path := synth.BlockPath(c)

	assert.False(t, r.Unregister(path))

	_, ok := r.MakeBlock(path)

	assert.False(t, ok)

	r.Register(func() synth.Block { return &TestBlock{} })

	c2, ok := r.MakeBlock(path)

	assert.True(t, ok)
	assert.NotNil(t, c2)
	assert.IsType(t, c, c2)
}
