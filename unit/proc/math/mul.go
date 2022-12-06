package math

// Mul adds two inputs.
type Mul struct {
	*BinaryOp
}

// NewMul makes a new Mul node.
func NewMul() *Mul {
	return &Mul{BinaryOp: NewBinaryOp()}
}

// Run performs the addition.
func (a *Mul) Run() {
	for i := 0; i < len(a.Out.BufferValues); i++ {
		a.Out.BufferValues[i] = a.BinaryOp.In1Buf[i] * a.BinaryOp.In2Buf[i]
	}
}
