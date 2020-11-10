package pow

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type Pow struct {
	inBuf  *node.Buffer
	expBuf *node.Buffer

	exp float64
}

func NewNode(in *node.Node, exp *node.Node) *node.Node {
	inputs := map[string]*node.Node{"In": in}
	controls := map[string]*node.Node{"Exp": exp}
	return node.New(New(), inputs, controls)
}

func New() *Pow {
	return &Pow{}
}

func (p *Pow) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	p.inBuf = node.GetOutput(inputs, "In")
	p.expBuf = node.GetOutput(controls, "Exp")
}

func (p *Pow) Configure() {
	p.exp = p.expBuf.Values[0]
}

func (p *Pow) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = math.Pow(p.inBuf.Values[i], p.exp)
	}
}
