package constant

import "github.com/jamestunnell/go-synth/node"

type Const struct {
	val float64
}

func NewNode(val float64) *node.Node {
	return node.New(New(val), map[string]*node.Node{}, map[string]*node.Node{})
}

func New(val float64) *Const {
	return &Const{val: val}
}

func (c *Const) Initialize(srate float64, inputs, controls map[string]*node.Node) {
}

func (c *Const) Configure() {
}

func (c *Const) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = c.val
	}
}
