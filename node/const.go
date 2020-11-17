package node

type Const struct {
	Value float64 `json:"value"`
}

func init() {
	WorkingRegistry().Register(&Const{Value: 0.0})
}

func NewConst(val float64) *Node {
	return NewNode(&Const{Value: val}, Map{}, Map{})
}

func (c *Const) Interface() *Interface {
	return &Interface{
		InputNames:      []string{},
		ControlDefaults: map[string]float64{},
	}
}

func (c *Const) Initialize(srate float64, inputs, controls Map) {
}

func (c *Const) Configure() {
}

func (c *Const) Run(out *Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = c.Value
	}
}
