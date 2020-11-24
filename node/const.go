package node

const ParamNameValue = "Value"

type Const struct {
	Value float64 `json:"value"`
}

func init() {
	WorkingRegistry().Register(&Const{Value: 0.0})
}

func NewConst(val float64) *Node {
	return &Node{
		Core:     &Const{},
		Params:   ParamMap{ParamNameValue: val},
		Controls: Map{},
		Inputs:   Map{},
	}
}

func (c *Const) Interface() *Interface {
	return &Interface{
		InputNames:      []string{},
		ControlDefaults: map[string]float64{},
		ParamTypes: map[string]ParamType{
			ParamNameValue: ParamTypeNumber,
		},
	}
}

func (c *Const) Initialize(args *InitArgs) error {
	c.Value = args.Params[ParamNameValue].(float64)

	return nil
}

func (c *Const) Configure() {
}

func (c *Const) Run(out *Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = c.Value
	}
}
