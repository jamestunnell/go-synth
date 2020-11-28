package decay

// State is for enumerating states
type State int

// StateMachine is an ADSR state machine.
type StateMachine struct {
	// state is the current SM state
	state State
	// triggerOn tracks if the trigger is active
	triggerOn bool
	// the current evelope output level
	level float64
	// decayMul is the multiplier to apply per sample to the envelope level
	decayMul float64
}

const (
	// Quiescent is the initial state
	Quiescent State = iota
	// Active state entered after trigger is detected completes
	Active
	// envStart is the envelope output immediately after a trigger is detected
	envStart = 1.0
	// QuiescentThreshold is the threshold below which the envelope
	// will drop to 0 and become quiescent
	QuiescentThreshold = 1e-12
)

// NewStateMachine makes a new StateMachine.
func NewStateMachine(srate, decayTime float64) *StateMachine {
	// sample periods to reach the target envelope output
	n := srate * decayTime

	return &StateMachine{
		state:     Quiescent,
		triggerOn: false,
		level:     0.0,
		decayMul:  Multiplier(int(n)),
	}
}

// State returns the current state.
func (sm *StateMachine) State() State {
	return sm.state
}

// Level returns the current state.
func (sm *StateMachine) Level() float64 {
	return sm.level
}

// Run executes the state machine and returns the current level (envelope output).
func (sm *StateMachine) Run(triggerVal float64) float64 {
	triggerOn := triggerVal > 0.0

	// detect trigger off-on transition
	if !sm.triggerOn && triggerOn {
		sm.level = envStart
		sm.state = Active
		sm.triggerOn = true

		return sm.level
	} else if sm.triggerOn && !triggerOn {
		sm.triggerOn = false
	}

	if sm.state == Active {
		if sm.level <= QuiescentThreshold {
			sm.state = Quiescent
			sm.level = 0.0
		} else {
			sm.level *= sm.decayMul
		}
	}

	return sm.level
}
