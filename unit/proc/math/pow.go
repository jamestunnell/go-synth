package math

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

// Pow raises input to a power.
type Pow struct {
	inBuf, expBuf *node.Buffer
	exp           float64
}

// ControlNameExp is the control name for the
// Pow exponent control
const ControlNameExp = "Exp"

// NewPow makes a new Pow instance.
func NewPow(in, exp *node.Node) *node.Node {
	return node.New(&Pow{},
		node.AddInput(InNameIn, in),
		node.AddControl(ControlNameExp, exp))
}

// Interface provides the node interface.
func (p *Pow) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.InputNames = []string{InNameIn}
	ifc.ControlDefaults[ControlNameExp] = 1.0

	return ifc
}

// Initialize initializes the node.
func (p *Pow) Initialize(args *node.InitArgs) error {
	p.inBuf = args.Inputs[InNameIn].Output()
	p.expBuf = args.Controls[ControlNameExp].Output()

	return nil
}

// Configure configures the node using latest output from the
// Exp control.
func (p *Pow) Configure() {
	p.exp = p.expBuf.Values[0]
}

// Run raises the node input to current exponent.
func (p *Pow) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = math.Pow(p.inBuf.Values[i], p.exp)
	}
}
