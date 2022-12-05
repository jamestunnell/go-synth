package math

import (
	"github.com/jamestunnell/go-synth"
)

// UnaryOp partially implements the node.Core interface
type UnaryOp struct {
	In *synth.TypedInput[float64]
}

// NewUnaryOp makes a new UnaryOp node that uses the given core.
func NewUnaryOp() *UnaryOp {
	return &UnaryOp{
		In: synth.NewFloat64Input(),
	}
}

// Initialize initializes the node.
func (u *UnaryOp) Initialize(srate float64, outDepth int) error {
	return nil
}

// Configure does nothing
func (u *UnaryOp) Configure() {
}
