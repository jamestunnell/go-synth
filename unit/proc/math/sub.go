package math

// Sub adds two inputs.
type Sub struct {
	*BinaryOp
}

// NewSub makes a new Sub node.
func NewSub() *Sub {
	return &Sub{BinaryOp: NewBinaryOp()}
}

// Run performs the addition.
func (a *Sub) Run() {
	for i := 0; i < len(a.Out.Buffer); i++ {
		a.Out.Buffer[i] = a.BinaryOp.In1.Output.Buffer[i] - a.BinaryOp.In2.Output.Buffer[i]
	}
}
