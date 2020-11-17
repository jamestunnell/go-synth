package mul

import "github.com/jamestunnell/go-synth/node"

type Mul struct {
	in1Buf, in2Buf *node.Buffer
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := node.Map{"In1": in1, "In2": in2}
	return node.NewNode(New(), inputs, node.Map{})
}

func New() *Mul {
	return &Mul{}
}

func (m *Mul) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{"In1", "In2"},
		ControlDefaults: map[string]float64{},
	}
}

func (m *Mul) Initialize(srate float64, inputs, controls node.Map) {
	m.in1Buf = inputs["In1"].Output()
	m.in2Buf = inputs["In2"].Output()
}

func (m *Mul) Configure() {
}

func (m *Mul) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = m.in1Buf.Values[i] * m.in2Buf.Values[i]
	}
}
