package sub_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/math/sub"
	"github.com/stretchr/testify/assert"
)

func TestAddXYHappyPath(t *testing.T) {
	in1 := array.OneShot([]float64{0.0, 0.1, 0.2})
	in2 := array.OneShot([]float64{-1.0, 0.5, 0.2})
	subXY := sub.XY(in1, in2)

	node.Initialize(subXY, 100.0, 3)
	node.Run(subXY)

	assert.Equal(t, 1.0, subXY.Buffer().Values[0])
	assert.Equal(t, -0.4, subXY.Buffer().Values[1])
	assert.Equal(t, 0.0, subXY.Buffer().Values[2])
}
