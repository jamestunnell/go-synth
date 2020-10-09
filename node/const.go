package node

type Const struct {
	val    float64
	outBuf *Buffer
}

func NewConst(val float64) Node {
	return &Const{val: val}
}

func (c *Const) Buffer() *Buffer {
	return c.outBuf
}

func (c *Const) Controls() map[string]Node {
	return map[string]Node{}
}

func (c *Const) Inputs() map[string]Node {
	return map[string]Node{}
}

func (c *Const) Initialize(srate float64, depth int) {
	c.outBuf = NewBuffer(depth)

	for i := 0; i < depth; i++ {
		c.outBuf.Values[i] = c.val
	}
}

func (c *Const) Configure() {
}

func (c *Const) Run() {
}
