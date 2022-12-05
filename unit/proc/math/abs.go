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
	abs := &Abs{}
	unaryOp := NewUnaryOp(abs)

	abs.UnaryOp = unaryOp

	return abs
}

// Run applies the absolute value
func (a *Abs) Run() {
	for i := 0; i < len(a.UnaryOp.Out.BufferValues); i++ {
		a.Out.BufferValues[i] = m.Abs(a.UnaryOp.InBuf[i])
	}
}
