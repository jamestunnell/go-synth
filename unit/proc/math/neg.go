package math

// Neg applies absolute value to an input.
type Neg struct {
	*UnaryOp
}

// NewNeg makes a new Neg block.
func NewNeg() *Neg {
	return &Neg{UnaryOp: NewUnaryOp()}
}

// Run applies the absolute value
func (a *Neg) Run() {
	for i := 0; i < len(a.UnaryOp.Out.BufferValues); i++ {
		a.Out.BufferValues[i] = -a.UnaryOp.InBuf[i]
	}
}
