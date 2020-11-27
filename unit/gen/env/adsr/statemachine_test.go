package adsr_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/env/adsr"
	"github.com/stretchr/testify/assert"
)

const (
	testSampleRate   = 10000.0
	testAttackMs     = 10
	testSamplePeriod = 1.0 / testSampleRate
	delta            = 1e-5
)

func TestSMStaysQuiescent(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    1.0,
		SustainLevel: 0.2,
		AttackMs:     5.0,
		DecayMs:      5.0,
		ReleaseMs:    5.0,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	assert.Equal(t, adsr.Quiescent, sm.State())

	// Make sure the state machine stays quiescent forever
	// when there is no trigger
	testSMHold(t, sm, 0.0, 1001, adsr.Quiescent, 0.0)
}

func TestSMTriggerLongEnoughToSustain(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    0.8,
		SustainLevel: 0.2,
		AttackMs:     4.0,
		DecayMs:      4.0,
		ReleaseMs:    4.0,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	// Positive trigger value should result in state change and should start
	// slewing right away
	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.2, 0.4, 0.6})

	// If the positive trigger continues, we should move into decay state and
	// start slewing downward toward sustain level
	testSMSlew(t, sm, 1.0, adsr.Decay, []float64{0.8, 0.65, 0.5, 0.35})

	// As long as the positive trigger continues we should stay in sustain state
	testSMHold(t, sm, 1.0, 300, adsr.Sustain, 0.2)

	// Once the trigger becomes non-positive, we should change state and start slewing to 0
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.15, 0.1, 0.05})

	// This last run would put the SM into quiescent state only if it reached (or passed) zero.
	level := sm.Run(0.0)

	assert.InDelta(t, 0.0, level, 1e-5)

	if level > 0.0 {
		assert.Equal(t, adsr.Release, sm.State())
	} else {
		assert.Equal(t, adsr.Quiescent, sm.State())
	}

	// As long as the non-positive trigger continues we should stay in quiescent state
	testSMHold(t, sm, 0.0, 300, adsr.Quiescent, 0.0)
}

func TestSMTriggerReleasedDuringAttack(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    0.8,
		SustainLevel: 0.2,
		AttackMs:     4.0,
		DecayMs:      4.0,
		ReleaseMs:    4.0,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.2, 0.4})

	// Release trigger - we should move to release state and slew down
	// at the normal release rate
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.35, 0.3, 0.25, 0.2, 0.15, 0.1, 0.05})

	// This last run would put the SM into quiescent state only if it reached (or passed) zero.
	level := sm.Run(0.0)

	assert.InDelta(t, 0.0, level, 1e-5)

	if level > 0.0 {
		assert.Equal(t, adsr.Release, sm.State())
	} else {
		assert.Equal(t, adsr.Quiescent, sm.State())
	}

	// As long as the non-positive trigger continues we should stay in quiescent state
	testSMHold(t, sm, 0.0, 300, adsr.Quiescent, 0.0)
}

func TestSMTriggerReleasedOutDuringDecay(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    0.8,
		SustainLevel: 0.2,
		AttackMs:     4.0,
		DecayMs:      4.0,
		ReleaseMs:    4.0,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.2, 0.4, 0.6})

	// hold trigger for part of the decay phase
	testSMSlew(t, sm, 1.0, adsr.Decay, []float64{0.8, 0.65, 0.5})

	// Release trigger - we should move to release state
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.45, 0.4, 0.35, 0.3, 0.25, 0.2, 0.15, 0.1, 0.05})

	// This last run would put the SM into quiescent state only if it reached (or passed) zero.
	level := sm.Run(0.0)

	assert.InDelta(t, 0.0, level, 1e-5)

	if level > 0.0 {
		assert.Equal(t, adsr.Release, sm.State())
	} else {
		assert.Equal(t, adsr.Quiescent, sm.State())
	}

	// As long as the non-positive trigger continues we should stay in quiescent state
	testSMHold(t, sm, 0.0, 300, adsr.Quiescent, 0.0)
}

func TestSMTriggerReactivatedBeforeQuiescent(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    0.8,
		SustainLevel: 0.2,
		AttackMs:     4.0,
		DecayMs:      4.0,
		ReleaseMs:    4.0,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.2, 0.4})
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.35, 0.3, 0.25, 0.2})
	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.4, 0.6})
}

func TestSMDecaySustainQuiescentAllStartMidwayThroughPeriod(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    1.0,
		SustainLevel: 0.25,
		AttackMs:     2.5,
		DecayMs:      3.0,
		ReleaseMs:    2.5,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.4, 0.8})
	testSMSlew(t, sm, 1.0, adsr.Decay, []float64{0.875, 0.625, 0.375})
	testSMHold(t, sm, 1.0, 50, adsr.Sustain, 0.25)
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.15, 0.05})
	testSMHold(t, sm, 0.0, 50, adsr.Quiescent, 0.0)
}

func TestSMDecayReleaseSoShortTheyGetSkipped(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    1.0,
		SustainLevel: 0.25,
		AttackMs:     2.5,
		DecayMs:      0.49,
		ReleaseMs:    0.99,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.4, 0.8})
	testSMHold(t, sm, 1.0, 50, adsr.Sustain, 0.25)
	testSMHold(t, sm, 0.0, 50, adsr.Quiescent, 0.0)
}

func TestSMAttackSoShortItGetsSkipped(t *testing.T) {
	p := &adsr.Params{
		PeakLevel:    1.0,
		SustainLevel: 0.25,
		AttackMs:     0.5,
		DecayMs:      3.0,
		ReleaseMs:    2.0,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Decay, []float64{0.875, 0.625, 0.375})
	testSMHold(t, sm, 1.0, 50, adsr.Sustain, 0.25)
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.125})
	testSMHold(t, sm, 0.0, 50, adsr.Quiescent, 0.0)
}

func testSMHold(
	t *testing.T,
	sm *adsr.StateMachine,
	trigger float64,
	n int,
	expectedState adsr.State,
	expectedLevel float64,
) bool {
	for i := 0; i < n; i++ {
		level := sm.Run(trigger)

		assert.Equal(t, expectedState, sm.State())
		assert.InDelta(t, expectedLevel, level, delta)
	}

	return !t.Failed()
}

func testSMSlew(
	t *testing.T,
	sm *adsr.StateMachine,
	trigger float64,
	expectedState adsr.State,
	expectedLevels []float64,
) bool {
	for _, expectedLevel := range expectedLevels {
		level := sm.Run(trigger)

		assert.Equal(t, expectedState, sm.State())
		assert.InDelta(t, expectedLevel, level, delta)
	}

	return !t.Failed()
}
