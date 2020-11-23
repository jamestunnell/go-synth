package math

import "github.com/jamestunnell/go-synth/node"

type Mul struct {
	*BinaryOp
}

func NewMul(in1, in2 *node.Node) *node.Node {
	return NewBinaryOp(&Mul{&BinaryOp{}}, in1, in2)
}

func (m *Mul) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = m.BinaryOp.In1Buf.Values[i] * m.BinaryOp.In2Buf.Values[i]
	}
}
