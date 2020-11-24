package nodetest

import (
	"errors"
	"math"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/util/param"
)

// TestCore is used for testing node.Node
type TestCore struct {
	InBuf, ControlBuf    *node.Buffer
	ControlVal, ParamVal float64
}

const (
	// InputName is the input for TestCore
	InputName = "In"
	// ParamName is the parm for TestCore
	ParamName = "Param"
	// ControlName is the control for TestCore
	ControlName = "Control"
	// ControlDefault is the default value for the control
	ControlDefault = 1.7
	// ParamType is the param type
	ParamType = param.Float
	// BadParamVal is for testing Initialize failure
	BadParamVal = math.Pi
)

// Interface provides the node interface.
func (tc *TestCore) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{InputName},
		ControlDefaults: map[string]float64{ControlName: ControlDefault},
		ParamTypes:      map[string]param.Type{ParamName: ParamType},
	}
}

// Initialize initializes the node.
// Returns a non-nil error if the param value equals BadParamVal.
func (tc *TestCore) Initialize(args *node.InitArgs) error {
	tc.InBuf = args.Inputs[InputName].Output()
	tc.ControlBuf = args.Controls[ControlName].Output()
	tc.ParamVal = args.Params[ParamName].Value().(float64)

	if tc.ParamVal == BadParamVal {
		return errors.New("hit the bad param val")
	}

	return nil
}

// Configure does nothing.
func (tc *TestCore) Configure() {
	tc.ControlVal = tc.ControlBuf.Values[0]
}

// Run just copies input to the given output buffer.
func (tc *TestCore) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = tc.InBuf.Values[i]
	}
}
