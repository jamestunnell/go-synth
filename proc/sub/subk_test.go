package sub_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/sub"
	"github.com/stretchr/testify/assert"
)

func TestAddKHappyPath(t *testing.T) {
	in := array.OneShot([]float64{0.0, 0.1, 0.2})
	subK := sub.K(in, 1.0)

	node.Initialize(subK, 100.0, 3)
	node.Run(subK)

	assert.Equal(t, -1.0, subK.Buffer().Values[0])
	assert.Equal(t, -0.9, subK.Buffer().Values[1])
	assert.Equal(t, -0.8, subK.Buffer().Values[2])
}
