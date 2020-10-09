package mul_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/math/mul"
	"github.com/stretchr/testify/assert"
)

func TestMulXYHappyPath(t *testing.T) {
	in1 := array.OneShot([]float64{0.0, 0.1, 0.2})
	in2 := array.OneShot([]float64{1.0, -1.0, 0.5})
	m := mul.New(in1, in2)

	node.Initialize(m, 100.0, 3)
	node.Run(m)

	assert.Equal(t, 0.0, m.Buffer().Values[0])
	assert.Equal(t, -0.1, m.Buffer().Values[1])
	assert.Equal(t, 0.1, m.Buffer().Values[2])
}