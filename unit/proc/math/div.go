package math

import "github.com/jamestunnell/go-synth/node"

type Div struct {
	*BinaryOp
}

func NewDiv(in1, in2 *node.Node) *node.Node {
	return NewBinaryOp(&Div{&BinaryOp{}}, in1, in2)
}

func (d *Div) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = d.BinaryOp.In1Buf.Values[i] / d.BinaryOp.In2Buf.Values[i]
	}
}
