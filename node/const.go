package node

import "github.com/jamestunnell/go-synth/util/param"

// ParamNameValue is the param name used for Const node
const ParamNameValue = "Value"

// Const is a node with constant output value
type Const struct {
	Value float64 `json:"value"`
}

func init() {
	WorkingRegistry().Register(&Const{Value: 0.0})
}

// NewConst makes a new Const node
func NewConst(val float64) *Node {
	addParam := AddParam(ParamNameValue, param.NewFloat(val))
	return New(&Const{}, addParam)
}

// Interface provides the node interface.
func (c *Const) Interface() *Interface {
	return &Interface{
		InputNames:      []string{},
		ControlDefaults: map[string]float64{},
		ParamTypes: map[string]param.Type{
			ParamNameValue: param.Float,
		},
	}
}

// Initialize initializes the node.
func (c *Const) Initialize(args *InitArgs) error {
	c.Value = args.Params[ParamNameValue].Value().(float64)

	return nil
}

// Configure does nothing
func (c *Const) Configure() {
}

// Run copies the const value to the given buffer.
func (c *Const) Run(out *Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = c.Value
	}
}
