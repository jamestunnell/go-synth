package white_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/unit/gen/noise/white"
)

func TestWhite(t *testing.T) {
	w := white.New()

	assert.True(t, w.Seed.SetValue(int64(12345)))

	w.Configure()
	w.Initialize(100.0, 1000)

	w.Run()

	vals := w.Out.Buffer().([]float64)

	for i := 0; i < len(vals); i++ {
		assert.GreaterOrEqual(t, vals[i], -1.0)
		assert.Less(t, vals[i], 1.0)
	}
}
