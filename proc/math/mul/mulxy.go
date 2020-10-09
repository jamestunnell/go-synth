package mul

import "github.com/jamestunnell/go-synth/node"

type mulXY struct {
	in1, in2 node.Node

	outBuf, in1Buf, in2Buf *node.Buffer
}

func XY(in1, in2 node.Node) node.Node {
	return &mulXY{in1: in1, in2: in2}
}

func (mul *mulXY) Buffer() *node.Buffer {
	return mul.outBuf
}

func (mul *mulXY) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (mul *mulXY) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In1": mul.in1,
		"In2": mul.in2,
	}
}

func (mul *mulXY) Initialize(srate float64, depth int) {
	mul.outBuf = node.NewBuffer(depth)
	mul.in1Buf = mul.in1.Buffer()
	mul.in2Buf = mul.in2.Buffer()
}

func (mul *mulXY) Configure() {
}

func (mul *mulXY) Run() {
	for i := 0; i < mul.outBuf.Length; i++ {
		mul.outBuf.Values[i] = mul.in1Buf.Values[i] * mul.in2Buf.Values[i]
	}
}
