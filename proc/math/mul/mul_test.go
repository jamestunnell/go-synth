package mul_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/proc/math/mul"
	"github.com/stretchr/testify/assert"
)

func TestMulHappyPath(t *testing.T) {
	in1 := oneshot.NewNode([]float64{0.0, 0.1, 0.2})
	in2 := oneshot.NewNode([]float64{1.0, -1.0, 0.5})
	n := mul.NewNode(in1, in2)

	n.Initialize(100.0, 3)
	n.Run()

	assert.Equal(t, 0.0, n.Output.Values[0])
	assert.Equal(t, -0.1, n.Output.Values[1])
	assert.Equal(t, 0.1, n.Output.Values[2])
}
