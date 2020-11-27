package adsr

import (
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/util/param"
)

// ADSR generates an ADSR envelope.
// Implements node.Core interface.
type ADSR struct {
	params       *Params
	stateMachine *StateMachine
	triggerBuff  *node.Buffer
}

const (
	// ParamNamePeakLevel is the name used for the peak level param
	ParamNamePeakLevel = "PeakLevel"
	// ParamNameSustainLevel is the name used for the sustain level param
	ParamNameSustainLevel = "SustainLevel"
	// ParamNameAttackMs is the name used for the attack time param
	ParamNameAttackMs = "AttackMs"
	// ParamNameDecayMs is the name used for the decay time param
	ParamNameDecayMs = "DecayMs"
	// ParamNameReleaseMs is the name used for the release time param
	ParamNameReleaseMs = "ReleaseMs"
	// InputNameTrigger is the name used for the trigger input
	InputNameTrigger = "Trigger"
)

// NewADSR makes a new ADSR node
func NewADSR(params *Params, mods ...node.Mod) *node.Node {
	mods = append(params.MakeMods(), mods...)

	return node.New(&ADSR{}, mods...)
}

// Interface provides the node interface
func (adsr *ADSR) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.ParamTypes = map[string]param.Type{
		ParamNamePeakLevel:    param.Float,
		ParamNameSustainLevel: param.Float,
		ParamNameAttackMs:     param.Float,
		ParamNameDecayMs:      param.Float,
		ParamNameReleaseMs:    param.Float,
	}

	ifc.InputNames = []string{InputNameTrigger}

	return ifc
}

// Initialize initializes the node, including making a new state machine.
// Returns a non-nil error if any of the params are invalid.
func (adsr *ADSR) Initialize(args *node.InitArgs) error {
	params := NewParamsFromMap(args.Params)

	if err := params.Validate(); err != nil {
		return err
	}

	adsr.triggerBuff = args.Inputs[InputNameTrigger].Output()
	adsr.params = params
	adsr.stateMachine = NewStateMachine(args.SampleRate, params)
	return nil
}

// Configure does nothing.
func (adsr *ADSR) Configure() {
}

// Run runs the state machine and places results in the given buffer.
func (adsr *ADSR) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = adsr.stateMachine.Run(adsr.triggerBuff.Values[i])
	}
}
