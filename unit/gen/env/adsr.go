package env

import (
	"fmt"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/env/adsr"
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

	stateMachine *adsr.StateMachine
}

// New makes a new ADSR node
func NewADSR() *ADSR {
	return &ADSR{
		AttackTime:   synth.NewFloat64Param(adsr.DefaultAttackTime),
		DecayTime:    synth.NewFloat64Param(adsr.DefaultDecayTime),
		SustainLevel: synth.NewFloat64Param(adsr.DefaultSustainLevel),
		ReleaseTime:  synth.NewFloat64Param(adsr.DefaultReleaseTime),
		Trigger:      synth.NewFloat64Input(),
		Out:          synth.NewFloat64Output(),
	}
}

// Initialize initializes the node, including making a new state machine.
// Returns a non-nil error if any of the params are invalid.
func (a *ADSR) Initialize(srate float64, outDepth int) error {
	a.Out.Initialize(outDepth)

	params := &adsr.Params{
		SustainLevel: a.SustainLevel.Value,
		AttackTime:   a.AttackTime.Value,
		DecayTime:    a.DecayTime.Value,
		ReleaseTime:  a.ReleaseTime.Value,
	}

	if err := params.Validate(); err != nil {
		return fmt.Errorf("invalid param(s): %w", err)
	}

	a.stateMachine = adsr.NewStateMachine(srate, params)

	return nil
}

// Configure does nothing.
func (a *ADSR) Configure() {
}

// Run runs the state machine and places results in the given buffer.
func (a *ADSR) Run() {
	for i := 0; i < len(a.Out.Buffer); i++ {
		a.Out.Buffer[i] = a.stateMachine.Run(a.Trigger.Output.Buffer[i])
	}
}
