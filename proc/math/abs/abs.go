package abs

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type absVal struct {
	in node.Node

	outBuf, inBuf *node.Buffer
}

func New(in node.Node) node.Node {
	return &absVal{in: in}
}

func (a *absVal) Buffer() *node.Buffer {
	return a.outBuf
}

func (a *absVal) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (a *absVal) Inputs() map[string]node.Node {
	return map[string]node.Node{"in": a.in}
}

func (a *absVal) Initialize(srate float64, depth int) {
	a.outBuf = node.NewBuffer(depth)
	a.inBuf = a.in.Buffer()
}

func (a *absVal) Configure() {
}

func (a *absVal) Run() {
	for i := 0; i < a.outBuf.Length; i++ {
		a.outBuf.Values[i] = math.Abs(a.inBuf.Values[i])
	}
}
