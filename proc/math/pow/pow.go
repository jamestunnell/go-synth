package pow

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type power struct {
	in  node.Node
	exp float64

	outBuf, inBuf *node.Buffer
}

func New(in node.Node, exp float64) node.Node {
	return &power{in: in, exp: exp}
}

func (p *power) Buffer() *node.Buffer {
	return p.outBuf
}

func (p *power) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (p *power) Inputs() map[string]node.Node {
	return map[string]node.Node{"in": p.in}
}

func (p *power) Initialize(srate float64, depth int) {
	p.outBuf = node.NewBuffer(depth)
	p.inBuf = p.in.Buffer()
}

func (p *power) Configure() {
}

func (p *power) Run() {
	for i := 0; i < p.outBuf.Length; i++ {
		p.outBuf.Values[i] = math.Pow(p.inBuf.Values[i], p.exp)
	}
}
