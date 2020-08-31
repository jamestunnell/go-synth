package processors

import (
	"github.com/jamestunnell/go-synth/node"
)

type AddXY struct {
	In1 interface{}
	In2 interface{}

	in1Buf *node.Buffer
	in2Buf *node.Buffer
}

func (a *AddXY) GetInterface() *node.Interface {
	return &node.Interface{
		Parameters: map[string]*node.ParamInfo{},
		Inputs:     []string{"In1", "In2"},
	}
}

func (a *AddXY) Initialize(srate float64) {
	a.in1Buf = a.In1.(*node.Node).Out
	a.in2Buf = a.In2.(*node.Node).Out
}

func (a *AddXY) Configure() {
}

func (a *AddXY) Sample(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = a.in1Buf.Values[i] + a.in2Buf.Values[i]
	}
}
