package math

import "github.com/jamestunnell/go-synth/node"

type Sub struct {
	*BinaryOp
}

func NewSub(in1, in2 *node.Node) *node.Node {
	return NewBinaryOp(&Sub{&BinaryOp{}}, in1, in2)
}

func (s *Sub) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = s.BinaryOp.In1Buf.Values[i] - s.BinaryOp.In2Buf.Values[i]
	}
}
