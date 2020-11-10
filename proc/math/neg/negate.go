package neg

import (
	"github.com/jamestunnell/go-synth/node"
)

type Neg struct {
	inBuf *node.Buffer
}

func NewNode(in *node.Node) *node.Node {
	inputs := map[string]*node.Node{"In": in}
	return node.New(New(), inputs, map[string]*node.Node{})
}

func New() *Neg {
	return &Neg{}
}

func (n *Neg) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	n.inBuf = node.GetOutput(inputs, "In")
}

func (n *Neg) Configure() {
}

func (n *Neg) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = -n.inBuf.Values[i]
	}
}
