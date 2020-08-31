package processors

import (
	"github.com/jamestunnell/go-synth/node"
)

type MulK struct {
	In interface{}
	K  interface{}

	k     float64
	kBuf  *node.Buffer
	inBuf *node.Buffer
}

func (m *MulK) GetInterface() *node.Interface {
	return &node.Interface{
		Parameters: map[string]*node.ParamInfo{
			"K": &node.ParamInfo{Required: false, Default: 0.0},
		},
		Inputs: []string{"In"},
	}
}

func (m *MulK) Initialize(srate float64) {
	m.kBuf = m.K.(*node.Node).Out
	m.inBuf = m.In.(*node.Node).Out
}

func (m *MulK) Configure() {
	m.k = m.kBuf.Values[0]
}

func (m *MulK) Sample(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = m.inBuf.Values[i] * m.k
	}
}
