package abs_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/unit/proc/math/abs"
	"github.com/stretchr/testify/assert"
)

func TestAbsHappyPath(t *testing.T) {
	in := oneshot.NewNode([]float64{1.0, 0.5, -0.5})
	n := abs.NewNode(in)

	n.Initialize(100.0, 3)
	n.Run()

	assert.Equal(t, 1.0, n.Output().Values[0])
	assert.Equal(t, 0.5, n.Output().Values[1])
	assert.Equal(t, 0.5, n.Output().Values[2])
}