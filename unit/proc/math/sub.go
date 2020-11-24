package math

import "github.com/jamestunnell/go-synth/node"

// Sub subtracts second input from first.
type Sub struct {
	*BinaryOp
}

// NewSub makes a new Sub node.
func NewSub(in1, in2 *node.Node) *node.Node {
	return NewBinaryOp(&Sub{&BinaryOp{}}, in1, in2)
}

// Run performs the subtraction.
func (s *Sub) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = s.BinaryOp.In1Buf.Values[i] - s.BinaryOp.In2Buf.Values[i]
	}
}
