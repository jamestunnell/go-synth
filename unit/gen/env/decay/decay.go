package decay

import (
	"fmt"

	"github.com/jamestunnell/go-synth"
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

	triggerBuf []float64
	outBuf     []float64
	sm         *StateMachine
}

const (
	DefaultDecayTime = 0.05
)

// New makes a new Decay node
func New() *Decay {
	d := &Decay{
		DecayTime: synth.NewFloat64Param(DefaultDecayTime),
		Trigger:   synth.NewFloat64Input(),
	}

	d.Out = synth.NewFloat64Output(d)

	return d
}

// Initialize initializes the node.
// Returns a non-nil error if the decay param is invalid.
func (d *Decay) Initialize(srate float64, outDepth int) error {
	decayTime := d.DecayTime.Value

	if decayTime <= 0.0 {
		return fmt.Errorf("decay time %f is not positive", decayTime)
	}

	d.sm = NewStateMachine(srate, decayTime)
	d.triggerBuf = d.Trigger.Output.Buffer().([]float64)
	d.outBuf = d.Out.Buffer().([]float64)

	return nil
}

// Configure does nothing.
func (d *Decay) Configure() {
}

// Run runs the exponential decay process.
func (d *Decay) Run() {
	for i := 0; i < len(d.outBuf); i++ {
		d.outBuf[i] = d.sm.Run(d.triggerBuf[i])
	}
}
