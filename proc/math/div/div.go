package div

import "github.com/jamestunnell/go-synth/node"

type Div struct {
	in1Buf, in2Buf *node.Buffer
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := map[string]*node.Node{"In1": in1, "In2": in2}
	return node.New(New(), inputs, map[string]*node.Node{})
}

func New() *Div {
	return &Div{}
}

func (d *Div) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	d.in1Buf = node.GetOutput(inputs, "In1")
	d.in2Buf = node.GetOutput(inputs, "In2")
}

func (d *Div) Configure() {
}

func (d *Div) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = d.in1Buf.Values[i] / d.in2Buf.Values[i]
	}
}
