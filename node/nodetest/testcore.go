package nodetest

import (
	"errors"
	"math"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/util/param"
)

type TestCore struct {
	InBuf, ControlBuf    *node.Buffer
	ControlVal, ParamVal float64
}

const (
	InputName   = "In"
	ParamName   = "Param"
	ControlName = "Control"

	ControlDefault = 1.7
	ParamType      = param.Float
	BadParamVal    = math.Pi
)

func (tc *TestCore) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{InputName},
		ControlDefaults: map[string]float64{ControlName: ControlDefault},
		ParamTypes:      map[string]param.Type{ParamName: ParamType},
	}
}

func (tc *TestCore) Initialize(args *node.InitArgs) error {
	tc.InBuf = args.Inputs[InputName].Output()
	tc.ControlBuf = args.Controls[ControlName].Output()
	tc.ParamVal = args.Params[ParamName].Value().(float64)

	if tc.ParamVal == BadParamVal {
		return errors.New("hit the bad param val")
	}

	return nil
}

func (tc *TestCore) Configure() {
	tc.ControlVal = tc.ControlBuf.Values[0]
}

func (tc *TestCore) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = tc.InBuf.Values[i]
	}
}
