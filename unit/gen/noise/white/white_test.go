package white_test

import (
	"math"
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/noise/white"
)

func TestWhite(t *testing.T) {
	n := white.New()

	n.Initialize(100.0, 1000)

	n.Run()

	vals := n.Output().Values

	for i := 0; i < len(vals); i++ {
		assert.GreaterOrEqual(t, vals[i], -1.0)
		assert.Less(t, vals[i], 1.0)
	}
}
