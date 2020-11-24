package math

import "github.com/jamestunnell/go-synth/node"

// Div divides first input by second input.
type Div struct {
	*BinaryOp
}

// NewDiv makes a new Div node
func NewDiv(in1, in2 *node.Node) *node.Node {
	return NewBinaryOp(&Div{&BinaryOp{}}, in1, in2)
}

// Run performs the division
func (d *Div) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = d.BinaryOp.In1Buf.Values[i] / d.BinaryOp.In2Buf.Values[i]
	}
}
