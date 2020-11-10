package mul

import "github.com/jamestunnell/go-synth/node"

type Mul struct {
	in1Buf, in2Buf *node.Buffer
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := map[string]*node.Node{"In1": in1, "In2": in2}
	return node.New(New(), inputs, map[string]*node.Node{})
}

func New() *Mul {
	return &Mul{}
}

func (m *Mul) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	m.in1Buf = node.GetOutput(inputs, "In1")
	m.in2Buf = node.GetOutput(inputs, "In2")
}

func (m *Mul) Configure() {
}

func (m *Mul) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = m.in1Buf.Values[i] * m.in2Buf.Values[i]
	}
}
