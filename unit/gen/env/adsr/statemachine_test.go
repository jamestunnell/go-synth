package adsr_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit/gen/env/adsr"
	"github.com/stretchr/testify/assert"
)

const delta = 1e-5

func TestSMStaysQuiescent(t *testing.T) {
	p := &adsr.Params{
		SustainLevel: 0.2,
		AttackTime:   0.005,
		DecayTime:    0.005,
		ReleaseTime:  0.005,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	assert.Equal(t, adsr.Quiescent, sm.State())

	testSMHold(t, sm, 0.0, 1001, adsr.Quiescent, 0.0)
}

func TestSMTriggerLongEnoughToSustain(t *testing.T) {
	p := &adsr.Params{
		SustainLevel: 0.2,
		AttackTime:   0.004,
		DecayTime:    0.004,
		ReleaseTime:  0.004,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.25, 0.5, 0.75})
	testSMSlew(t, sm, 1.0, adsr.Decay, []float64{1.0, 0.8, 0.6, 0.4})
	testSMHold(t, sm, 1.0, 300, adsr.Sustain, 0.2)
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.15, 0.1, 0.05})
	testSMHold(t, sm, 0.0, 300, adsr.Quiescent, 0.0)
}

func TestSMTriggerReleasedDuringAttack(t *testing.T) {
	p := &adsr.Params{
		SustainLevel: 0.2,
		AttackTime:   0.004,
		DecayTime:    0.004,
		ReleaseTime:  0.004,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.25, 0.5})
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.45, 0.4, 0.35, 0.3, 0.25, 0.2, 0.15, 0.1, 0.05})
	testSMHold(t, sm, 0.0, 300, adsr.Quiescent, 0.0)
}

func TestSMTriggerReleasedDuringDecay(t *testing.T) {
	p := &adsr.Params{
		SustainLevel: 0.2,
		AttackTime:   0.004,
		DecayTime:    0.004,
		ReleaseTime:  0.004,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.25, 0.5, 0.75})
	testSMSlew(t, sm, 1.0, adsr.Decay, []float64{1.0, 0.8, 0.6, 0.4})
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.35, 0.3, 0.25, 0.2, 0.15, 0.1, 0.05})
	testSMHold(t, sm, 0.0, 300, adsr.Quiescent, 0.0)
}

func TestSMTriggerReactivatedBeforeQuiescent(t *testing.T) {
	p := &adsr.Params{
		SustainLevel: 0.2,
		AttackTime:   0.004,
		DecayTime:    0.004,
		ReleaseTime:  0.004,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.25, 0.5})
	testSMSlew(t, sm, 0.0, adsr.Release, []float64{0.45, 0.4, 0.35, 0.3, 0.25, 0.2})
	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.45, 0.7})
}

func TestSMDecaySustainQuiescentAllStartMidwayThroughPeriod(t *testing.T) {
	p := &adsr.Params{
		SustainLevel: 0.25,
		AttackTime:   0.0025,
		DecayTime:    0.003,
		ReleaseTime:  0.0025,
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
		SustainLevel: 0.25,
		AttackTime:   0.0025,
		DecayTime:    0.00049,
		ReleaseTime:  0.00099,
	}
	sm := adsr.NewStateMachine(1000.0, p)

	testSMSlew(t, sm, 1.0, adsr.Attack, []float64{0.4, 0.8})
	testSMHold(t, sm, 1.0, 50, adsr.Sustain, 0.25)
	testSMHold(t, sm, 0.0, 50, adsr.Quiescent, 0.0)
}

func TestSMAttackSoShortItGetsSkipped(t *testing.T) {
	p := &adsr.Params{
		SustainLevel: 0.25,
		AttackTime:   0.0005,
		DecayTime:    0.003,
		ReleaseTime:  0.002,
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
