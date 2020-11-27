package adsr

// State is for enumerating states
type State int

// StateMachine is an ADSR state machine.
type StateMachine struct {
	state                              State
	triggerOn                          bool
	level, peakLevel, sustainLevel     float64
	attackSlew, decaySlew, releaseSlew float64
}

const (
	// Quiescent is the initial state
	Quiescent State = iota
	// Attack state entered after a trigger is detected
	Attack
	// Decay state entered after attack completes
	Decay
	// Sustain state entered after decay completes
	Sustain
	// Release state entered after a trigger is released
	Release
)

// NewStateMachine makes a new StateMachine.
func NewStateMachine(srate float64, p *Params) *StateMachine {
	periodsPerMs := srate / 1000.0

	return &StateMachine{
		state:        Quiescent,
		triggerOn:    false,
		level:        0.0,
		peakLevel:    p.PeakLevel,
		sustainLevel: p.SustainLevel,
		attackSlew:   p.PeakLevel / (p.AttackMs * periodsPerMs),
		decaySlew:    (p.PeakLevel - p.SustainLevel) / (p.DecayMs * periodsPerMs),
		releaseSlew:  p.SustainLevel / (p.ReleaseMs * periodsPerMs),
	}
}

// State returns the current state.
func (sm *StateMachine) State() State {
	return sm.state
}

// Run executes the state machine and returns the current level (envelope output).
func (sm *StateMachine) Run(triggerVal float64) float64 {
	triggerOn := triggerVal > 0.0

	// detect trigger on-off and off-on transition
	if triggerOn != sm.triggerOn {
		if sm.triggerOn { // trigger transition from on to off
			sm.state = Release
		} else { // trigger transition from off to on
			sm.state = Attack
		}

		sm.triggerOn = triggerOn
	}

	switch sm.state {
	case Attack:
		sm.attack()
	case Decay:
		sm.decay()
	case Release:
		sm.release()
	}

	return sm.level
}

func (sm *StateMachine) attack() {
	sm.level += sm.attackSlew

	if sm.level == sm.peakLevel {
		sm.state = Decay
	}
	if sm.level > sm.peakLevel {
		over := sm.level - sm.peakLevel
		adjustedLevel := sm.peakLevel - (1-over/sm.attackSlew)*sm.decaySlew

		if adjustedLevel > sm.sustainLevel {
			sm.level = adjustedLevel
			sm.state = Decay
		} else {
			sm.level = sm.sustainLevel
			sm.state = Sustain
		}
	}
}

func (sm *StateMachine) decay() {
	sm.level -= sm.decaySlew
	if sm.level <= sm.sustainLevel {
		sm.level = sm.sustainLevel
		sm.state = Sustain
	}
}

func (sm *StateMachine) release() {
	sm.level -= sm.releaseSlew
	if sm.level <= 0.0 {
		sm.level = 0.0
		sm.state = Quiescent
	}
}
