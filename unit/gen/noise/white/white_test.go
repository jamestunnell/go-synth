package white_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/noise/white"
	"github.com/stretchr/testify/assert"
)

func TestWhite(t *testing.T) {
	w := white.New()

	w.Initialize(100.0, 10)

	w.Run()

	for i := 0; i < w.Output().Length; i++ {
		assert.GreaterOrEqual(t, w.Output().Values[i], -1.0)
		assert.LessOrEqual(t, w.Output().Values[i], 1.0)
	}
}
