package div_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/proc/math/div"
	"github.com/stretchr/testify/assert"
)

func TestDivHappyPath(t *testing.T) {
	in1 := oneshot.NewNode([]float64{0.0, 0.1, 0.2, 1.0})
	in2 := oneshot.NewNode([]float64{1.0, -1.0, 0.5, 2.0})
	n := div.NewNode(in1, in2)

	n.Initialize(100.0, 4)
	n.Run()

	assert.Equal(t, 0.0, n.Output.Values[0])
	assert.Equal(t, -0.1, n.Output.Values[1])
	assert.Equal(t, 0.4, n.Output.Values[2])
	assert.Equal(t, 0.5, n.Output.Values[3])
}
