package math

import (
	m "math"

	"github.com/jamestunnell/go-synth/node"
)

type Abs struct {
	*UnaryOp
}

func NewAbs(in *node.Node) *node.Node {
	return NewUnaryOp(&Abs{&UnaryOp{}}, in)
}

func (a *Abs) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = m.Abs(a.UnaryOp.InBuf.Values[i])
	}
}
