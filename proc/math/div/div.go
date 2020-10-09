package div

import "github.com/jamestunnell/go-synth/node"

type Div struct {
	in1, in2 node.Node

	outBuf, in1Buf, in2Buf *node.Buffer
}

func New(in1, in2 node.Node) node.Node {
	return &Div{in1: in1, in2: in2}
}

func (d *Div) Buffer() *node.Buffer {
	return d.outBuf
}

func (d *Div) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (d *Div) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In1": d.in1,
		"In2": d.in2,
	}
}

func (d *Div) Initialize(srate float64, depth int) {
	d.outBuf = node.NewBuffer(depth)
	d.in1Buf = d.in1.Buffer()
	d.in2Buf = d.in2.Buffer()
}

func (d *Div) Configure() {
}

func (d *Div) Run() {
	for i := 0; i < d.outBuf.Length; i++ {
		d.outBuf.Values[i] = d.in1Buf.Values[i] / d.in2Buf.Values[i]
	}
}
