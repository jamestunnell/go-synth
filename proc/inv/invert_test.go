package inv_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/inv"
	"github.com/stretchr/testify/assert"
)

func TestInvertHappyPath(t *testing.T) {
	in := array.OneShot([]float64{1.0, 0.5, -0.5})
	invert := inv.Invert(in)

	node.Initialize(invert, 100.0, 3)
	node.Run(invert)

	assert.Equal(t, 1.0, invert.Buffer().Values[0])
	assert.Equal(t, 2.0, invert.Buffer().Values[1])
	assert.Equal(t, -2.0, invert.Buffer().Values[2])
}
