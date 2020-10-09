package mul_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/generators/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/processors/mul"
	"github.com/stretchr/testify/assert"
)

func TestMulKHappyPath(t *testing.T) {
	in := array.OneShot([]float64{0.0, 0.2, -0.3})
	mulK := mul.K(in, 2.0)

	node.Initialize(mulK, 100.0, 3)
	node.Run(mulK)

	assert.Equal(t, 0.0, mulK.Buffer().Values[0])
	assert.Equal(t, 0.4, mulK.Buffer().Values[1])
	assert.Equal(t, -0.6, mulK.Buffer().Values[2])
}
