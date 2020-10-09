package add

import "github.com/jamestunnell/go-synth/node"

type Add struct {
	in1, in2 node.Node

	outBuf, in1Buf, in2Buf *node.Buffer
}

func New(in1, in2 node.Node) node.Node {
	return &Add{in1: in1, in2: in2}
}

func (a *Add) Buffer() *node.Buffer {
	return a.outBuf
}

func (a *Add) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (a *Add) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In1": a.in1,
		"In2": a.in2,
	}
}

func (a *Add) Initialize(srate float64, depth int) {
	a.outBuf = node.NewBuffer(depth)
	a.in1Buf = a.in1.Buffer()
	a.in2Buf = a.in2.Buffer()
}

func (a *Add) Configure() {
}

func (a *Add) Run() {
	for i := 0; i < a.outBuf.Length; i++ {
		a.outBuf.Values[i] = a.in1Buf.Values[i] + a.in2Buf.Values[i]
	}
}
