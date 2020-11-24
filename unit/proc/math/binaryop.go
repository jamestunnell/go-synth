package math

import "github.com/jamestunnell/go-synth/node"

// BinaryOp partially implements the node.Core interface.
type BinaryOp struct {
	In1Buf, In2Buf *node.Buffer
}

const (
	// InNameIn1 is the input name for the first input of a binary op node
	InNameIn1 = "In1"
	// InNameIn2 is the input name for the second input of a binary op node
	InNameIn2 = "In2"
)

// NewBinaryOp makes a new binary op node that uses the given core.
func NewBinaryOp(c node.Core, in1, in2 *node.Node) *node.Node {
	return node.New(c,
		node.AddInput(InNameIn1, in1),
		node.AddInput(InNameIn2, in2))
}

// Interface provides the node interface.
func (b *BinaryOp) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.InputNames = []string{InNameIn1, InNameIn2}

	return ifc
}

// Initialize initializes the node.
func (b *BinaryOp) Initialize(args *node.InitArgs) error {
	b.In1Buf = args.Inputs[InNameIn1].Output()
	b.In2Buf = args.Inputs[InNameIn2].Output()

	return nil
}

// Configure does nothing
func (b *BinaryOp) Configure() {
}
