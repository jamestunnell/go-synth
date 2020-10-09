package abs_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/math/abs"
	"github.com/stretchr/testify/assert"
)

func TestInvertHappyPath(t *testing.T) {
	in := array.OneShot([]float64{1.0, 0.5, -0.5})
	a := abs.New(in)

	node.Initialize(a, 100.0, 3)
	node.Run(a)

	assert.Equal(t, 1.0, a.Buffer().Values[0])
	assert.Equal(t, 0.5, a.Buffer().Values[1])
	assert.Equal(t, 0.5, a.Buffer().Values[2])
}
