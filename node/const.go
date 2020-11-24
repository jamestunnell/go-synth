package node

import "github.com/jamestunnell/go-synth/util/param"

const ParamNameValue = "Value"

type Const struct {
	Value float64 `json:"value"`
}

func init() {
	WorkingRegistry().Register(&Const{Value: 0.0})
}

func NewConst(val float64) *Node {
	addParam := MakeAddParam(ParamNameValue, param.NewFloat(val))
	return New(&Const{}, addParam)
}

func (c *Const) Interface() *Interface {
	return &Interface{
		InputNames:      []string{},
		ControlDefaults: map[string]float64{},
		ParamTypes: map[string]param.Type{
			ParamNameValue: param.Float,
		},
	}
}

func (c *Const) Initialize(args *InitArgs) error {
	c.Value = args.Params[ParamNameValue].Value().(float64)

	return nil
}

func (c *Const) Configure() {
}

func (c *Const) Run(out *Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = c.Value
	}
}
