package sub

import "github.com/jamestunnell/go-synth/node"

type Sub struct {
	in1, in2 node.Node

	outBuf, in1Buf, in2Buf *node.Buffer
}

func New(in1, in2 node.Node) node.Node {
	return &Sub{in1: in1, in2: in2}
}

func (s *Sub) Buffer() *node.Buffer {
	return s.outBuf
}

func (s *Sub) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (s *Sub) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In1": s.in1,
		"In2": s.in2,
	}
}

func (s *Sub) Initialize(srate float64, depth int) {
	s.outBuf = node.NewBuffer(depth)
	s.in1Buf = s.in1.Buffer()
	s.in2Buf = s.in2.Buffer()
}

func (s *Sub) Configure() {
}

func (s *Sub) Run() {
	for i := 0; i < s.outBuf.Length; i++ {
		s.outBuf.Values[i] = s.in1Buf.Values[i] - s.in2Buf.Values[i]
	}
}
