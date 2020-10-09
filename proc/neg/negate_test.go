package neg_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/neg"
	"github.com/stretchr/testify/assert"
)

func TestInvertHappyPath(t *testing.T) {
	in := array.OneShot([]float64{1.0, 0.5, -0.5})
	negate := neg.New(in)

	node.Initialize(negate, 100.0, 3)
	node.Run(negate)

	assert.Equal(t, -1.0, negate.Buffer().Values[0])
	assert.Equal(t, -0.5, negate.Buffer().Values[1])
	assert.Equal(t, 0.5, negate.Buffer().Values[2])
}
