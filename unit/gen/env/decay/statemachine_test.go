package decay_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/env/decay"
	"github.com/stretchr/testify/assert"
)

const delta = 1e-5

func TestSMStaysQuiescent(t *testing.T) {
	sm := decay.NewStateMachine(1000.0, 0.01)

	assert.Equal(t, decay.Quiescent, sm.State())

	testSMHold(t, sm, 0.0, 1001, decay.Quiescent, 0.0)
}

func TestSMStaysActiveAfterTriggerRelease(t *testing.T) {
	sm := decay.NewStateMachine(1000.0, 0.02)
	mul := decay.Multiplier(int(1000.0 * 0.02))

	level := sm.Run(1.0)

	assert.Equal(t, decay.Active, sm.State())
	assert.InDelta(t, 1.0, level, delta)

	for i := 0; i < 20; i++ {
		newLevel := sm.Run(0.0)

		assert.Equal(t, decay.Active, sm.State())
		assert.InDelta(t, level*mul, newLevel, delta)

		level = newLevel
	}

	assert.InDelta(t, decay.TargetDecay, sm.Level(), delta)
}

func TestSMDecayContinuesUntilQuiescentThreshold(t *testing.T) {
	sm := decay.NewStateMachine(1000.0, 0.02)
	mul := decay.Multiplier(int(1000.0 * 0.02))
	level := sm.Run(1.0)

	assert.Equal(t, decay.Active, sm.State())
	assert.InDelta(t, 1.0, level, delta)

	for level > decay.QuiescentThreshold {
		newLevel := sm.Run(0.0)

		assert.Equal(t, decay.Active, sm.State())
		assert.InDelta(t, level*mul, newLevel, delta)

		level = newLevel
	}

	level = sm.Run(0.0)

	assert.Equal(t, decay.Quiescent, sm.State())
	assert.Equal(t, 0.0, level)
}

func TestSMCanBeRetriggeredAfterRelease(t *testing.T) {
	sm := decay.NewStateMachine(1000.0, 0.02)
	mul := decay.Multiplier(int(1000.0 * 0.02))
	level1 := sm.Run(1.0)

	assert.Equal(t, decay.Active, sm.State())
	assert.InDelta(t, 1.0, level1, delta)

	level2 := sm.Run(0.0)

	assert.Equal(t, decay.Active, sm.State())
	assert.InDelta(t, level1*mul, level2, delta)

	level3 := sm.Run(1.0)

	assert.Equal(t, decay.Active, sm.State())
	assert.InDelta(t, 1.0, level3, delta)

	level4 := sm.Run(1.0)

	assert.Equal(t, decay.Active, sm.State())
	assert.InDelta(t, level3*mul, level4, delta)
}

func testSMHold(
	t *testing.T,
	sm *decay.StateMachine,
	trigger float64,
	n int,
	expectedState decay.State,
	expectedLevel float64,
) bool {
	for i := 0; i < n; i++ {
		level := sm.Run(trigger)

		assert.Equal(t, expectedState, sm.State())
		assert.InDelta(t, expectedLevel, level, delta)
	}

	return !t.Failed()
}
