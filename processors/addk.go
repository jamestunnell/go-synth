package processors

import (
	"github.com/jamestunnell/go-synth/node"
)

type AddK struct {
	In interface{}
	K  interface{}

	k     float64
	kBuf  *node.Buffer
	inBuf *node.Buffer
}

func (a *AddK) GetInterface() *node.Interface {
	return &node.Interface{
		Parameters: map[string]*node.ParamInfo{
			"K": &node.ParamInfo{Required: false, Default: 0.0},
		},
		Inputs: []string{"In"},
	}
}

func (a *AddK) Initialize(srate float64) {
	a.kBuf = a.K.(*node.Node).Out
	a.inBuf = a.In.(*node.Node).Out
}

func (a *AddK) Configure() {
	a.k = a.kBuf.Values[0]
}

func (a *AddK) Sample(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = a.inBuf.Values[i] + a.k
	}
}
