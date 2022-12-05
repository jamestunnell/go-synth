package adsr

import (
	"fmt"

	"github.com/jamestunnell/go-synth"
)

// ADSR generates an ADSR envelope.
// Implements node.Core interface.
type ADSR struct {
	SustainLevel *synth.TypedParam[float64]
	AttackTime   *synth.TypedParam[float64]
	DecayTime    *synth.TypedParam[float64]
	ReleaseTime  *synth.TypedParam[float64]

	Trigger *synth.TypedInput[float64]
	Out     *synth.TypedOutput[float64]

	stateMachine *StateMachine
	triggerBuf   []float64
}

const (
	DefaultSustainLevel = 0.5
	DefaultAttackTime   = 0.04
	DefaultDecayTime    = 0.03
	DefaultReleaseTime  = 0.02
)

// New makes a new ADSR node
func New() *ADSR {
	adsr := &ADSR{
		AttackTime:   synth.NewFloat64Param(DefaultAttackTime),
		DecayTime:    synth.NewFloat64Param(DefaultDecayTime),
		SustainLevel: synth.NewFloat64Param(DefaultSustainLevel),
		ReleaseTime:  synth.NewFloat64Param(DefaultReleaseTime),
		Trigger:      synth.NewFloat64Input(),
	}

	adsr.Out = synth.NewFloat64Output(adsr)

	return adsr
}

// Initialize initializes the node, including making a new state machine.
// Returns a non-nil error if any of the params are invalid.
func (adsr *ADSR) Initialize(srate float64, outDepth int) error {
	adsr.Out.Initialize(outDepth)

	params := &Params{
		SustainLevel: adsr.SustainLevel.Value,
		AttackTime:   adsr.AttackTime.Value,
		DecayTime:    adsr.DecayTime.Value,
		ReleaseTime:  adsr.ReleaseTime.Value,
	}

	if err := params.Validate(); err != nil {
		return fmt.Errorf("invalid param(s): %w", err)
	}

	adsr.triggerBuf = adsr.Trigger.Output.Buffer().([]float64)
	adsr.stateMachine = NewStateMachine(srate, params)

	return nil
}

// Configure does nothing.
func (adsr *ADSR) Configure() {
}

// Run runs the state machine and places results in the given buffer.
func (adsr *ADSR) Run() {
	for i := 0; i < len(adsr.Out.BufferValues); i++ {
		adsr.Out.BufferValues[i] = adsr.stateMachine.Run(adsr.triggerBuf[i])
	}
}
