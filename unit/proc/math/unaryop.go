package math

import (
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/mod"
)

// UnaryOp partially implements the node.Core interface
type UnaryOp struct {
	InBuf *node.Buffer
}

const (
	// InNameIn is the input name for a unary op core
	InNameIn = "In"
)

// NewUnaryOp makes a new UnaryOp node that uses the given core.
func NewUnaryOp(c node.Core, in *node.Node) *node.Node {
	return node.New(c, mod.Input(InNameIn, in))
}

// Interface provides the node interface.
func (u *UnaryOp) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.InputNames = []string{InNameIn}

	return ifc
}

// Initialize initializes the node.
func (u *UnaryOp) Initialize(args *node.InitArgs) error {
	u.InBuf = args.Inputs[InNameIn].Output()

	return nil
}

// Configure does nothing
func (u *UnaryOp) Configure() {
}
