package sub

import "github.com/jamestunnell/go-synth/node"

type subXY struct {
	in1, in2 node.Node

	outBuf, in1Buf, in2Buf *node.Buffer
}

func XY(in1, in2 node.Node) node.Node {
	return &subXY{in1: in1, in2: in2}
}

func (sub *subXY) Buffer() *node.Buffer {
	return sub.outBuf
}

func (sub *subXY) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (sub *subXY) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In1": sub.in1,
		"In2": sub.in2,
	}
}

func (sub *subXY) Initialize(srate float64, depth int) {
	sub.outBuf = node.NewBuffer(depth)
	sub.in1Buf = sub.in1.Buffer()
	sub.in2Buf = sub.in2.Buffer()
}

func (sub *subXY) Configure() {
}

func (sub *subXY) Run() {
	for i := 0; i < sub.outBuf.Length; i++ {
		sub.outBuf.Values[i] = sub.in1Buf.Values[i] - sub.in2Buf.Values[i]
	}
}
