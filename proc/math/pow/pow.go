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

func init() {
	node.WorkingRegistry().RegisterCore(New())
}

func NewNode(in *node.Node, exp *node.Node) *node.Node {
	inputs := node.Map{"In": in}
	controls := node.Map{"Exp": exp}
	return node.NewNode(New(), inputs, controls)
}

func New() *Pow {
	return &Pow{}
}

func (p *Pow) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{"In"},
		ControlDefaults: map[string]float64{"Exp": 1.0},
	}
}

func (p *Pow) Initialize(srate float64, inputs, controls node.Map) {
	p.inBuf = inputs["In"].Output()
	p.expBuf = controls["Exp"].Output()
}

func (p *Pow) Configure() {
	p.exp = p.expBuf.Values[0]
}

func (p *Pow) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = math.Pow(p.inBuf.Values[i], p.exp)
	}
}
