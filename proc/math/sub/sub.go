package sub

import "github.com/jamestunnell/go-synth/node"

type Sub struct {
	in1Buf, in2Buf *node.Buffer
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := node.Map{"In1": in1, "In2": in2}
	return node.NewNode(New(), inputs, node.Map{})
}

func New() *Sub {
	return &Sub{}
}

func (s *Sub) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{"In1", "In2"},
		ControlDefaults: map[string]float64{},
	}
}

func (s *Sub) Initialize(srate float64, inputs, controls node.Map) {
	s.in1Buf = inputs["In1"].Output()
	s.in2Buf = inputs["In2"].Output()
}

func (s *Sub) Configure() {
}

func (s *Sub) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = s.in1Buf.Values[i] - s.in2Buf.Values[i]
	}
}
