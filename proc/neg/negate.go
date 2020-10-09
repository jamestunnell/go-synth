package neg

import (
	"github.com/jamestunnell/go-synth/node"
)

type negate struct {
	in node.Node

	outBuf, inBuf *node.Buffer
}

func New(in node.Node) node.Node {
	return &negate{in: in}
}

func (neg *negate) Buffer() *node.Buffer {
	return neg.outBuf
}

func (neg *negate) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (neg *negate) Inputs() map[string]node.Node {
	return map[string]node.Node{"in": neg.in}
}

func (neg *negate) Initialize(srate float64, depth int) {
	neg.outBuf = node.NewBuffer(depth)
	neg.inBuf = neg.in.Buffer()
}

func (neg *negate) Configure() {
}

func (neg *negate) Run() {
	for i := 0; i < neg.outBuf.Length; i++ {
		neg.outBuf.Values[i] = -neg.inBuf.Values[i]
	}
}
