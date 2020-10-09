package add_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/gen/array"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/add"
	"github.com/stretchr/testify/assert"
)

func TestAddKHappyPath(t *testing.T) {
	in := array.OneShot([]float64{0.0, 0.1, 0.2})
	addK := add.K(in, 1.0)

	node.Initialize(addK, 100.0, 3)
	node.Run(addK)

	assert.Equal(t, 1.0, addK.Buffer().Values[0])
	assert.Equal(t, 1.1, addK.Buffer().Values[1])
	assert.Equal(t, 1.2, addK.Buffer().Values[2])
}
