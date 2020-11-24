package math

import "github.com/jamestunnell/go-synth/node"

type BinaryOp struct {
	In1Buf, In2Buf *node.Buffer
}

const (
	InNameIn1 = "In1"
	InNameIn2 = "In2"
)

func NewBinaryOp(c node.Core, in1, in2 *node.Node) *node.Node {
	return node.New(c,
		node.MakeAddInput(InNameIn1, in1),
		node.MakeAddInput(InNameIn2, in2))
}

func (b *BinaryOp) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.InputNames = []string{InNameIn1, InNameIn2}

	return ifc
}

func (b *BinaryOp) Initialize(args *node.InitArgs) error {
	b.In1Buf = args.Inputs[InNameIn1].Output()
	b.In2Buf = args.Inputs[InNameIn2].Output()

	return nil
}

func (b *BinaryOp) Configure() {
}
