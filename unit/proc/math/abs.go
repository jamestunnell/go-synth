package math

import (
	m "math"

	"github.com/jamestunnell/go-synth"
)

// Abs applies absolute value to an input.
type Abs struct {
	In  *synth.TypedInput[float64]
	Out *synth.TypedOutput[float64]

	inBuf []float64
}

// NewAbs makes a new Abs block.
func NewAbs() *Abs {
	a := &Abs{
		In: synth.NewFloat64Input(),
	}

	a.Out = synth.NewFloat64Output(a)

	return a
}

// Initialize initializes the block.
func (a *Abs) Initialize(srate float64, outDepth int) error {
	a.Out.Initialize(outDepth)

	a.inBuf = a.In.Output.Buffer().([]float64)

	return nil
}

// Configure does nothing
func (a *Abs) Configure() {
}

// Run applies the absolute value
func (a *Abs) Run() {
	for i := 0; i < len(a.Out.BufferValues); i++ {
		a.Out.BufferValues[i] = m.Abs(a.inBuf[i])
	}
}
