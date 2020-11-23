package math

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type Pow struct {
	inBuf, expBuf *node.Buffer
	exp           float64
}

const ControlNameExp = "Exp"

func NewPow(in, exp *node.Node) *node.Node {
	return &node.Node{
		Core:     &Pow{},
		Inputs:   node.Map{InNameIn: in},
		Controls: node.Map{ControlNameExp: exp},
		Params:   node.ParamMap{},
	}
}

func (p *Pow) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.InputNames = []string{InNameIn}
	ifc.ControlDefaults[ControlNameExp] = 1.0

	return ifc
}

func (p *Pow) Initialize(args *node.InitArgs) error {
	p.inBuf = args.Inputs[InNameIn].Output()
	p.expBuf = args.Controls[ControlNameExp].Output()

	return nil
}

func (p *Pow) Configure() {
	p.exp = p.expBuf.Values[0]
}

func (p *Pow) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = math.Pow(p.inBuf.Values[i], p.exp)
	}
}
