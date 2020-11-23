package math

import "github.com/jamestunnell/go-synth/node"

type UnaryOp struct {
	InBuf *node.Buffer
}

const (
	InNameIn = "In"
)

func NewUnaryOp(c node.Core, in *node.Node) *node.Node {
	return &node.Node{
		Core:     c,
		Inputs:   node.Map{InNameIn: in},
		Controls: node.Map{},
		Params:   node.ParamMap{},
	}
}

func (u *UnaryOp) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.InputNames = []string{InNameIn}

	return ifc
}

func (u *UnaryOp) Initialize(args *node.InitArgs) error {
	u.InBuf = args.Inputs[InNameIn].Output()

	return nil
}

func (u *UnaryOp) Configure() {
}
