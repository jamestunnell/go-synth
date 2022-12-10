package math

// Div adds two inputs.
type Div struct {
	*BinaryOp
}

// NewDiv makes a new Div node.
func NewDiv() *Div {
	return &Div{BinaryOp: NewBinaryOp()}
}

// Run performs the addition.
func (a *Div) Run() {
	for i := 0; i < len(a.Out.Buffer); i++ {
		a.Out.Buffer[i] = a.BinaryOp.In1.Output.Buffer[i] / a.BinaryOp.In2.Output.Buffer[i]
	}
}
