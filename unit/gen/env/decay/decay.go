package decay

import (
	"fmt"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/mod"
	"github.com/jamestunnell/go-synth/util/param"
)

// Decay generates an exponential decay envelope.
// The decay rate is controlled by the decay time parameter, and is set
// so the envelope decays from 1 down to 0.01 (-40 dB) when decay time
// has passed since the trigger was detected.
// Implements node.Core interface.
type Decay struct {
	triggerBuf *node.Buffer
	sm         *StateMachine
}

const (
	// ParamNameDecayTime is the name used for the decay time param
	ParamNameDecayTime = "DecayTime"
	// InputNameTrigger is the name used for the trigger input
	InputNameTrigger = "Trigger"
)

// NewDecay makes a new Decay node
func NewDecay(decayTime float64, moreMods ...node.Mod) *node.Node {
	mods := []node.Mod{mod.Param(ParamNameDecayTime, param.NewFloat(decayTime))}

	return node.New(&Decay{}, append(mods, moreMods...)...)
}

// Interface provides the node interface
func (d *Decay) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.ParamTypes = map[string]param.Type{
		ParamNameDecayTime: param.Float,
	}

	ifc.InputNames = []string{InputNameTrigger}

	return ifc
}

// Initialize initializes the node.
// Returns a non-nil error if the decay param is invalid.
func (d *Decay) Initialize(args *node.InitArgs) error {
	decayTime := args.Params[ParamNameDecayTime].Value().(float64)

	if decayTime <= 0.0 {
		return fmt.Errorf("decay time %f is not positive", decayTime)
	}

	d.sm = NewStateMachine(args.SampleRate, decayTime)
	d.triggerBuf = args.Inputs[InputNameTrigger].Output()

	return nil
}

// Configure does nothing.
func (d *Decay) Configure() {
}

// Run runs the exponential decay process.
func (d *Decay) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = d.sm.Run(d.triggerBuf.Values[i])
	}
}
