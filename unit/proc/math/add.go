package math

// Add adds two inputs.
type Add struct {
	*BinaryOp
}

// NewAdd makes a new Add node.
func NewAdd() *Add {
	return &Add{BinaryOp: NewBinaryOp()}
}

// Run performs the addition.
func (a *Add) Run() {
	for i := 0; i < len(a.Out.BufferValues); i++ {
		a.Out.BufferValues[i] = a.BinaryOp.In1Buf[i] + a.BinaryOp.In2Buf[i]
	}
}
