package math

import (
	"math"

	"github.com/jamestunnell/go-synth"
)

// Pow raises input to a power.
type Pow struct {
	In  *synth.TypedInput[float64]
	Exp *synth.TypedControl[float64]
	Out *synth.TypedOutput[float64]

	exp float64
}

// NewPow makes a new Pow block.
func NewPow() *Pow {
	return &Pow{
		In:  synth.NewFloat64Input(),
		Exp: synth.NewFloat64Control(1.0),
		Out: synth.NewFloat64Output(),
	}
}

// Initialize initializes the block.
func (p *Pow) Initialize(srate float64, outDepth int) error {
	p.Out.Initialize(outDepth)

	return nil
}

// Configure does nothing.
func (p *Pow) Configure() {
	p.exp = p.Exp.Output.Buffer[0]
}

// Run raises the input to the exponent.
func (p *Pow) Run() {
	for i := 0; i < len(p.Out.Buffer); i++ {
		p.Out.Buffer[i] = math.Pow(p.In.Output.Buffer[i], p.exp)
	}
}
