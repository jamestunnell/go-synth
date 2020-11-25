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

// Mul makes a function that modifies a node so the output is
// multiplied by the given mul node (before any output adding).
func Mul(mul *node.Node) node.Mod {
	return func(n *node.Node) {
		n.Mul = mul
	}
}

// MulK makes a function that modifies a node so the output is
// multiplied by the given mul value (before any output adding).
func MulK(mul float64) node.Mod {
	return func(n *node.Node) {
		n.Mul = node.NewK(mul)
	}
}

// Add makes a function that modifies a node so the output is added
// with the given add node (after any output multiplication).
func Add(add *node.Node) node.Mod {
	return func(n *node.Node) {
		n.Add = add
	}
}

// AddK makes a function that modifies a node so the output is added
// with the given add value (after any output multiplication).
func AddK(add float64) node.Mod {
	return func(n *node.Node) {
		n.Add = node.NewK(add)
	}
}
