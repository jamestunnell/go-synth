package processors

import (
	"github.com/jamestunnell/go-synth/node"
)

type MulXY struct {
	In1 interface{}
	In2 interface{}

	in1Buf *node.Buffer
	in2Buf *node.Buffer
}

func (m *MulXY) GetInterface() *node.Interface {
	return &node.Interface{
		Parameters: map[string]*node.ParamInfo{},
		Inputs:     []string{"In1", "In2"},
	}
}

func (m *MulXY) Initialize(srate float64) {
	m.in1Buf = m.In1.(*node.Node).Out
	m.in2Buf = m.In2.(*node.Node).Out
}

func (m *MulXY) Configure() {
}

func (m *MulXY) Sample(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = m.in1Buf.Values[i] * m.in2Buf.Values[i]
	}
}
