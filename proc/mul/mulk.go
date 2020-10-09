package mul

import "github.com/jamestunnell/go-synth/node"

type mulK struct {
	in  node.Node
	val float64

	outBuf, inBuf *node.Buffer
}

func K(in node.Node, val float64) node.Node {
	return &mulK{in: in, val: val}
}

func (mul *mulK) Buffer() *node.Buffer {
	return mul.outBuf
}

func (mul *mulK) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (mul *mulK) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In": mul.in,
	}
}

func (mul *mulK) Initialize(srate float64, depth int) {
	mul.outBuf = node.NewBuffer(depth)
	mul.inBuf = mul.in.Buffer()
}

func (mul *mulK) Configure() {
}

func (mul *mulK) Run() {
	for i := 0; i < mul.outBuf.Length; i++ {
		mul.outBuf.Values[i] = mul.inBuf.Values[i] * mul.val
	}
}
