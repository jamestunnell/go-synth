package constant

import "github.com/jamestunnell/go-synth/node"

type Const struct {
	val    float64
	outBuf *node.Buffer
}

func New(val float64) node.Node {
	return &Const{val: val}
}

func (c *Const) Buffer() *node.Buffer {
	return c.outBuf
}

func (c *Const) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (c *Const) Inputs() map[string]node.Node {
	return map[string]node.Node{}
}

func (c *Const) Initialize(srate float64, depth int) {
	c.outBuf = node.NewBuffer(depth)

	for i := 0; i < depth; i++ {
		c.outBuf.Values[i] = c.val
	}
}

func (c *Const) Configure() {
}

func (c *Const) Run() {
}
