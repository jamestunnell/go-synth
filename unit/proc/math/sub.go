package math

// Sub adds two inputs.
type Sub struct {
	*BinaryOp
}

// NewSub makes a new Sub node.
func NewSub() *Sub {
	add := &Sub{}
	binaryOp := NewBinaryOp(add)

	add.BinaryOp = binaryOp

	return add
}

// Run performs the addition.
func (a *Sub) Run() {
	for i := 0; i < len(a.Out.BufferValues); i++ {
		a.Out.BufferValues[i] = a.BinaryOp.In1Buf[i] - a.BinaryOp.In2Buf[i]
	}
}
