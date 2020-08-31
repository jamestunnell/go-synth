package processors

import "github.com/jamestunnell/go-synth/node"

var (
	BuiltinProcessors = []node.Core{
		&AddK{},
		&AddXY{},
		&Invert{},
		&MulK{},
		&MulXY{},
	}
)
