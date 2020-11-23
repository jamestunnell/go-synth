package math

import "github.com/jamestunnell/go-synth/node"

type Add struct {
	*BinaryOp
}

func NewAdd(in1, in2 *node.Node) *node.Node {
	return NewBinaryOp(&Add{&BinaryOp{}}, in1, in2)
}

func (a *Add) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = a.BinaryOp.In1Buf.Values[i] + a.BinaryOp.In2Buf.Values[i]
	}
}
