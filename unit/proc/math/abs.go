package math

import (
	m "math"
)

// Abs applies absolute value to an input.
type Abs struct {
	*UnaryOp
}

// NewAbs makes a new Abs block.
func NewAbs() *Abs {
	return &Abs{UnaryOp: NewUnaryOp()}
}

// Run applies the absolute value
func (a *Abs) Run() {
	for i := 0; i < len(a.UnaryOp.Out.Buffer); i++ {
		a.Out.Buffer[i] = m.Abs(a.UnaryOp.In.Output.Buffer[i])
	}
}
