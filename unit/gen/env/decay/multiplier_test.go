package decay_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/env/decay"
	"github.com/stretchr/testify/assert"
)

func TestMultiplierN1(t *testing.T) {
	assert.InDelta(t, decay.TargetDecay, decay.Multiplier(1), delta)
}

func TestMultiplierN2(t *testing.T) {
	decayMul := decay.Multiplier(2)
	result := 1.0 * decayMul * decayMul

	assert.InDelta(t, decay.TargetDecay, result, delta)
}

func TestMultiplierN17(t *testing.T) {
	decayMul := decay.Multiplier(17)

	result := 1.0
	for i := 0; i < 17; i++ {
		result *= decayMul
	}

	assert.InDelta(t, decay.TargetDecay, result, delta)
}
