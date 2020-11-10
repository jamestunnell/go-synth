package sub

import "github.com/jamestunnell/go-synth/node"

type Sub struct {
	in1Buf, in2Buf *node.Buffer
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := map[string]*node.Node{"In1": in1, "In2": in2}
	return node.New(New(), inputs, map[string]*node.Node{})
}

func New() *Sub {
	return &Sub{}
}

func (s *Sub) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	s.in1Buf = node.GetOutput(inputs, "In1")
	s.in2Buf = node.GetOutput(inputs, "In2")
}

func (s *Sub) Configure() {
}

func (s *Sub) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = s.in1Buf.Values[i] - s.in2Buf.Values[i]
	}
}
