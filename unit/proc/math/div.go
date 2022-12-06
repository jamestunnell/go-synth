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
	for i := 0; i < len(a.Out.BufferValues); i++ {
		a.Out.BufferValues[i] = a.BinaryOp.In1Buf[i] / a.BinaryOp.In2Buf[i]
	}
}
