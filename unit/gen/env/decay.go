package env

import (
	"fmt"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/env/decay"
)

// Decay generates an exponential decay envelope.
// The decay rate is controlled by the decay time parameter, and is set
// so the envelope decays from 1 down to 0.01 (-40 dB) when decay time
// has passed since the trigger was detected.
// Implements node.Core interface.
type Decay struct {
	DecayTime *synth.TypedParam[float64]
	Trigger   *synth.TypedInput[float64]
	Out       *synth.TypedOutput[float64]

	sm *decay.StateMachine
}

// New makes a new Decay node
func NewDecay() *Decay {
	return &Decay{
		DecayTime: synth.NewFloat64Param(decay.DefaultDecayTime),
		Trigger:   synth.NewFloat64Input(),
		Out:       synth.NewFloat64Output(),
	}
}

// Initialize initializes the node.
// Returns a non-nil error if the decay param is invalid.
func (d *Decay) Initialize(srate float64, outDepth int) error {
	decayTime := d.DecayTime.Value

	if decayTime <= 0.0 {
		return fmt.Errorf("decay time %f is not positive", decayTime)
	}

	d.sm = decay.NewStateMachine(srate, decayTime)

	return nil
}

// Configure does nothing.
func (d *Decay) Configure() {
}

// Run runs the exponential decay process.
func (d *Decay) Run() {
	for i := 0; i < len(d.Out.Buffer); i++ {
		d.Out.Buffer[i] = d.sm.Run(d.Trigger.Output.Buffer[i])
	}
}
