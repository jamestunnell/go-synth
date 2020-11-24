package node

import "github.com/jamestunnell/go-synth/util/param"

// AddControl makes a function that adds the given control.
func AddControl(name string, control *Node) Mod {
	return func(n *Node) {
		n.Controls[name] = control
	}
}

// AddInput makes a function that adds the given input.
func AddInput(name string, input *Node) Mod {
	return func(n *Node) {
		n.Inputs[name] = input
	}
}

// AddParam makes a function that adds the given param.
func AddParam(name string, p *param.Param) Mod {
	return func(n *Node) {
		n.Params[name] = p
	}
}
