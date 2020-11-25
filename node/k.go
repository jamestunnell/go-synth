package node

import "github.com/jamestunnell/go-synth/util/param"

// ParamNameValue is the param name used for K node
const ParamNameValue = "Value"

// K is a node with constant output value
type K struct {
	Value float64 `json:"value"`
}

func init() {
	WorkingRegistry().Register(&K{})
}

// NewK makes a new K (const) node
func NewK(val float64) *Node {
	return New(&K{}, func(n *Node) {
		n.Params[ParamNameValue] = param.NewFloat(val)
	})
}

// Interface provides the node interface.
func (k *K) Interface() *Interface {
	return &Interface{
		InputNames:      []string{},
		ControlDefaults: map[string]float64{},
		ParamTypes: map[string]param.Type{
			ParamNameValue: param.Float,
		},
	}
}

// Initialize initializes the node.
func (k *K) Initialize(args *InitArgs) error {
	k.Value = args.Params[ParamNameValue].Value().(float64)

	return nil
}

// Configure does nothing
func (k *K) Configure() {
}

// Run copies the const value to the given buffer.
func (k *K) Run(out *Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = k.Value
	}
}
