package math

import (
	"github.com/jamestunnell/go-synth"
)

// BinaryOp partially implements the synth.Block interface.
type BinaryOp struct {
	In1, In2 *synth.TypedInput[float64]
	Out      *synth.TypedOutput[float64]

	In1Buf, In2Buf []float64
}

// NewBinaryOp makes a new BinaryOp which can be used to make a binary math block.
func NewBinaryOp(parent synth.Block) *BinaryOp {
	return &BinaryOp{
		In1: synth.NewFloat64Input(),
		In2: synth.NewFloat64Input(),
		Out: synth.NewFloat64Output(parent),
	}
}

// Initialize initializes the block.
func (b *BinaryOp) Initialize(srate float64, outDepth int) error {
	b.Out.Initialize(outDepth)

	b.In1Buf = b.In1.Output.Buffer().([]float64)
	b.In2Buf = b.In2.Output.Buffer().([]float64)

	return nil
}

// Configure does nothing
func (b *BinaryOp) Configure() {
}
