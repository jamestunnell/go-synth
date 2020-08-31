package processors

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type Invert struct {
	In interface{}

	inBuf *node.Buffer
}

func (inv *Invert) GetInterface() *node.Interface {
	return &node.Interface{
		Parameters: map[string]*node.ParamInfo{},
		Inputs:     []string{"In"},
	}
}

func (inv *Invert) Initialize(srate float64) {
	inv.inBuf = inv.In.(*node.Node).Out
}

func (inv *Invert) Configure() {
}

func (inv *Invert) Sample(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		x := inv.inBuf.Values[i]

		var y float64
		if x == 0.0 {
			y = math.MaxFloat64
		} else {
			y = 1.0 / x
		}

		out.Values[i] = y
	}
}
