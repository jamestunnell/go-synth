package math

import (
	"github.com/jamestunnell/go-synth"
)

// UnaryOp partially implements the synth.Block interface
type UnaryOp struct {
	In  *synth.TypedInput[float64]
	Out *synth.TypedOutput[float64]

	InBuf []float64
}

// NewUnaryOp makes a new UnaryOp which can be used to make a unary math block.
func NewUnaryOp(parent synth.Block) *UnaryOp {
	return &UnaryOp{
		In:  synth.NewFloat64Input(),
		Out: synth.NewFloat64Output(parent),
	}
}

// Initialize initializes the node.
func (u *UnaryOp) Initialize(srate float64, outDepth int) error {
	u.Out.Initialize(outDepth)

	u.InBuf = u.In.Output.Buffer().([]float64)

	return nil
}

// Configure does nothing
func (u *UnaryOp) Configure() {
}
