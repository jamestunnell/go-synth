package math

import (
	"math"

	"github.com/jamestunnell/go-synth"
)

// Pow raises input to a power.
type Pow struct {
	In  *synth.TypedInput[float64]
	Exp *synth.TypedParam[float64]
	Out *synth.TypedOutput[float64]

	inBuf []float64
	exp   float64
}

// NewPow makes a new Pow block.
func NewPow() *Pow {
	pow := &Pow{
		In:  synth.NewFloat64Input(),
		Exp: synth.NewFloat64Param(1),
	}

	pow.Out = synth.NewFloat64Output(pow)

	return pow
}

// Initialize initializes the block.
func (p *Pow) Initialize(srate float64, outDepth int) error {
	p.Out.Initialize(outDepth)

	p.inBuf = p.In.Output.Buffer().([]float64)
	p.exp = p.Exp.GetValue().(float64)

	return nil
}

// Configure does nothing.
func (p *Pow) Configure() {
}

// Run raises the input to the exponent.
func (p *Pow) Run() {
	for i := 0; i < len(p.Out.BufferValues); i++ {
		p.Out.BufferValues[i] = math.Pow(p.inBuf[i], p.exp)
	}
}
