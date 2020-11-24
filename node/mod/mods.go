package mod

import (
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/util/param"
)

// Control makes a function that adds/modifies a control.
func Control(name string, control *node.Node) node.Mod {
	return func(n *node.Node) {
		n.Controls[name] = control
	}
}

// Input makes a function that adds/modifies an input.
func Input(name string, input *node.Node) node.Mod {
	return func(n *node.Node) {
		n.Inputs[name] = input
	}
}

// Param makes a function that adds/modifies a param.
func Param(name string, p *param.Param) node.Mod {
	return func(n *node.Node) {
		n.Params[name] = p
	}
}
