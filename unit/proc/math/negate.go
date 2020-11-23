package math

import (
	"github.com/jamestunnell/go-synth/node"
)

type Neg struct {
	*UnaryOp
}

func NewNeg(in *node.Node) *node.Node {
	return NewUnaryOp(&Neg{&UnaryOp{}}, in)
}

func (n *Neg) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = -n.UnaryOp.InBuf.Values[i]
	}
}
