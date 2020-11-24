package math

import (
	"github.com/jamestunnell/go-synth/node"
)

// Neg negates the input
type Neg struct {
	*UnaryOp
}

// NewNeg makes a new Neg node.
func NewNeg(in *node.Node) *node.Node {
	return NewUnaryOp(&Neg{&UnaryOp{}}, in)
}

// Run perorms the negation.
func (n *Neg) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = -n.UnaryOp.InBuf.Values[i]
	}
}
