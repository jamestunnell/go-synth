package math

import (
	m "math"

	"github.com/jamestunnell/go-synth/node"
)

// Abs applies absolute value to an input.
type Abs struct {
	*UnaryOp
}

// NewAbs makes a new Abs node
func NewAbs(in *node.Node) *node.Node {
	return NewUnaryOp(&Abs{&UnaryOp{}}, in)
}

// Run applies the absolute value
func (a *Abs) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = m.Abs(a.UnaryOp.InBuf.Values[i])
	}
}
