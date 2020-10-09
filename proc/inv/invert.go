package inv

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type invert struct {
	in node.Node

	outBuf, inBuf *node.Buffer
}

func New(in node.Node) node.Node {
	return &invert{in: in}
}

func (inv *invert) Buffer() *node.Buffer {
	return inv.outBuf
}

func (inv *invert) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (inv *invert) Inputs() map[string]node.Node {
	return map[string]node.Node{"in": inv.in}
}

func (inv *invert) Initialize(srate float64, depth int) {
	inv.outBuf = node.NewBuffer(depth)
	inv.inBuf = inv.in.Buffer()
}

func (inv *invert) Configure() {
}

func (inv *invert) Run() {
	for i := 0; i < inv.outBuf.Length; i++ {
		x := inv.inBuf.Values[i]

		var y float64
		if x == 0.0 {
			y = math.MaxFloat64
		} else {
			y = 1.0 / x
		}

		inv.outBuf.Values[i] = y
	}
}
