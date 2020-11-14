package neg

import (
	"github.com/jamestunnell/go-synth/node"
)

type Neg struct {
	inBuf *node.Buffer
}

func init() {
	node.WorkingRegistry().RegisterCore(New())
}

func NewNode(in *node.Node) *node.Node {
	inputs := node.Map{"In": in}
	return node.NewNode(New(), inputs, node.Map{})
}

func New() *Neg {
	return &Neg{}
}

func (n *Neg) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{"In"},
		ControlDefaults: map[string]float64{},
	}
}

func (n *Neg) Initialize(srate float64, inputs, controls node.Map) {
	n.inBuf = inputs["In"].Output()
}

func (n *Neg) Configure() {
}

func (n *Neg) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = -n.inBuf.Values[i]
	}
}
