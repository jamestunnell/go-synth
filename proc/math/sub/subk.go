package sub

import "github.com/jamestunnell/go-synth/node"

type subK struct {
	in  node.Node
	val float64

	outBuf, inBuf *node.Buffer
}

func K(in node.Node, val float64) node.Node {
	return &subK{in: in, val: val}
}

func (sub *subK) Buffer() *node.Buffer {
	return sub.outBuf
}

func (sub *subK) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (sub *subK) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In": sub.in,
	}
}

func (sub *subK) Initialize(srate float64, depth int) {
	sub.outBuf = node.NewBuffer(depth)
	sub.inBuf = sub.in.Buffer()
}

func (sub *subK) Configure() {
}

func (sub *subK) Run() {
	for i := 0; i < sub.outBuf.Length; i++ {
		sub.outBuf.Values[i] = sub.inBuf.Values[i] - sub.val
	}
}
