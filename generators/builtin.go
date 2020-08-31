package generators

import "github.com/jamestunnell/go-synth/node"

var (
	BuiltinGenerators = []node.Core{
		&Array{},
		&Osc{},
	}
)
