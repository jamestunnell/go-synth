package math

import (
	"github.com/jamestunnell/go-synth"
)

// Neg applies absolute value to an input.
type Neg struct {
	In  *synth.TypedInput[float64]
	Out *synth.TypedOutput[float64]

	inBuf []float64
}

// NewNeg makes a new Neg block.
func NewNeg() *Neg {
	n := &Neg{
		In: synth.NewFloat64Input(),
	}

	n.Out = synth.NewFloat64Output(n)

	return n
}

// Initialize initializes the block.
func (n *Neg) Initialize(srate float64, outDepth int) error {
	n.Out.Initialize(outDepth)

	n.inBuf = n.In.Output.Buffer().([]float64)

	return nil
}

// Configure does nothing
func (n *Neg) Configure() {
}

// Run applies the absolute value
func (n *Neg) Run() {
	for i := 0; i < len(n.Out.BufferValues); i++ {
		n.Out.BufferValues[i] = -n.inBuf[i]
	}
}
