package adsr

// State is for enumerating states
type State int

// StateMachine is an ADSR state machine.
type StateMachine struct {
	state                                  State
	triggerOn                              bool
	level, sustainLevel, sustainTransition float64
	attackSlew, decaySlew, releaseSlew     float64
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

	peakLevel           = 1.0
	closeEnough         = 1e-12
	decayTransition     = peakLevel - closeEnough
	quiescentTransition = closeEnough
)

// NewStateMachine makes a new StateMachine.
func NewStateMachine(srate float64, p *Params) *StateMachine {
	return &StateMachine{
		state:             Quiescent,
		triggerOn:         false,
		level:             0.0,
		sustainLevel:      p.SustainLevel,
		sustainTransition: p.SustainLevel + closeEnough,
		attackSlew:        peakLevel / (p.AttackTime * srate),
		decaySlew:         (peakLevel - p.SustainLevel) / (p.DecayTime * srate),
		releaseSlew:       p.SustainLevel / (p.ReleaseTime * srate),
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

	if sm.level > peakLevel {
		over := sm.level - peakLevel
		adjustedLevel := peakLevel - (1-over/sm.attackSlew)*sm.decaySlew

		if adjustedLevel > sm.sustainLevel {
			sm.level = adjustedLevel
			sm.state = Decay
		} else {
			sm.level = sm.sustainLevel
			sm.state = Sustain
		}
	} else if sm.level > decayTransition {
		sm.state = Decay
		sm.level = peakLevel
	}
}

func (sm *StateMachine) decay() {
	sm.level -= sm.decaySlew
	if sm.level <= sm.sustainTransition {
		sm.level = sm.sustainLevel
		sm.state = Sustain
	}
}

func (sm *StateMachine) release() {
	sm.level -= sm.releaseSlew
	if sm.level <= quiescentTransition {
		sm.level = 0.0
		sm.state = Quiescent
	}
}
