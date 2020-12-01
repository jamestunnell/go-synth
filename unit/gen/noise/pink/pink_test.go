package pink_test

import (
	"testing"
	"time"

	"github.com/jamestunnell/go-synth/unit/gen/noise/pink"
	"github.com/stretchr/testify/assert"
)

func TestPink(t *testing.T) {
	n := pink.NewFromSeed(time.Now().UTC().UnixNano())

	n.Initialize(100.0, 1000)

	vals := n.Output().Values

	for i := 0; i < len(vals); i++ {
		assert.GreaterOrEqual(t, vals[i], -1.0)
		assert.Less(t, vals[i], 1.0)
	}
}
