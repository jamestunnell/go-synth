package add

import "github.com/jamestunnell/go-synth/node"

type addK struct {
	in  node.Node
	val float64

	outBuf, inBuf *node.Buffer
}

func K(in node.Node, val float64) node.Node {
	return &addK{in: in, val: val}
}

func (add *addK) Buffer() *node.Buffer {
	return add.outBuf
}

func (add *addK) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (add *addK) Inputs() map[string]node.Node {
	return map[string]node.Node{
		"In": add.in,
	}
}

func (add *addK) Initialize(srate float64, depth int) {
	add.outBuf = node.NewBuffer(depth)
	add.inBuf = add.in.Buffer()
}

func (add *addK) Configure() {
}

func (add *addK) Run() {
	for i := 0; i < add.outBuf.Length; i++ {
		add.outBuf.Values[i] = add.inBuf.Values[i] + add.val
	}
}
