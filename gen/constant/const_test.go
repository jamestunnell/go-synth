package constant_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/constant"
	"github.com/jamestunnell/go-synth/node"
	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	const val = 2.5

	c := constant.New(val)

	node.Initialize(c, 100.0, 3)
	node.Run(c)

	assert.Equal(t, val, c.Buffer().Values[0])
	assert.Equal(t, val, c.Buffer().Values[1])
	assert.Equal(t, val, c.Buffer().Values[2])
}
