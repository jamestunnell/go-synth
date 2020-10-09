package add

import "github.com/jamestunnell/go-synth/node"

type addXY struct {
	in1, in2 node.Node

	outBuf, in1Buf, in2Buf *node.Buffer
}

func XY(in1, in2 node.Node) node.Node {
	return &addXY{in1: in1, in2: in2}
}

func (add *addXY) Buffer() *node.Buffer {
	return add.outBuf
}

func (add *addXY) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (add *addXY) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In1": add.in1,
		"In2": add.in2,
	}
}

func (add *addXY) Initialize(srate float64, depth int) {
	add.outBuf = node.NewBuffer(depth)
	add.in1Buf = add.in1.Buffer()
	add.in2Buf = add.in2.Buffer()
}

func (add *addXY) Configure() {
}

func (add *addXY) Run() {
	for i := 0; i < add.outBuf.Length; i++ {
		add.outBuf.Values[i] = add.in1Buf.Values[i] + add.in2Buf.Values[i]
	}
}
