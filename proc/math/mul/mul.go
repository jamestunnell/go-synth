package mul

import "github.com/jamestunnell/go-synth/node"

type Mul struct {
	in1, in2 node.Node

	outBuf, in1Buf, in2Buf *node.Buffer
}

func New(in1, in2 node.Node) node.Node {
	return &Mul{in1: in1, in2: in2}
}

func (m *Mul) Buffer() *node.Buffer {
	return m.outBuf
}

func (m *Mul) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (m *Mul) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In1": m.in1,
		"In2": m.in2,
	}
}

func (m *Mul) Initialize(srate float64, depth int) {
	m.outBuf = node.NewBuffer(depth)
	m.in1Buf = m.in1.Buffer()
	m.in2Buf = m.in2.Buffer()
}

func (m *Mul) Configure() {
}

func (m *Mul) Run() {
	for i := 0; i < m.outBuf.Length; i++ {
		m.outBuf.Values[i] = m.in1Buf.Values[i] * m.in2Buf.Values[i]
	}
}
