package brown_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/noise/brown"
	"github.com/stretchr/testify/assert"
)

func TestBrown(t *testing.T) {
	n := brown.New()

	n.Initialize(100.0, 2000)

	n.Run()

	for i := 0; i < n.Output().Length; i++ {
		assert.GreaterOrEqual(t, n.Output().Values[i], -1.0)
		assert.LessOrEqual(t, n.Output().Values[i], 1.0)
	}
}
