package neg_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/proc/math/neg"
	"github.com/stretchr/testify/assert"
)

func TestInvertHappyPath(t *testing.T) {
	in := oneshot.NewNode([]float64{1.0, 0.5, -0.5})
	n := neg.NewNode(in)

	n.Initialize(100.0, 3)
	n.Run()

	assert.Equal(t, -1.0, n.Output.Values[0])
	assert.Equal(t, -0.5, n.Output.Values[1])
	assert.Equal(t, 0.5, n.Output.Values[2])
}
