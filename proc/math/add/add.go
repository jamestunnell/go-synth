package add

import "github.com/jamestunnell/go-synth/node"

type Add struct {
	in1Buf, in2Buf *node.Buffer
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := map[string]*node.Node{"In1": in1, "In2": in2}
	return node.New(New(), inputs, map[string]*node.Node{})
}

func New() *Add {
	return &Add{}
}

func (a *Add) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	a.in1Buf = node.GetOutput(inputs, "In1")
	a.in2Buf = node.GetOutput(inputs, "In2")
}

func (a *Add) Configure() {
}

func (a *Add) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = a.in1Buf.Values[i] + a.in2Buf.Values[i]
	}
}
