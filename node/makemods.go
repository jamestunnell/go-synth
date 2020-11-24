package node

import "github.com/jamestunnell/go-synth/util/param"

func MakeAddControl(name string, control *Node) Mod {
	return func(n *Node) {
		n.Controls[name] = control
	}
}

func MakeAddInput(name string, input *Node) Mod {
	return func(n *Node) {
		n.Inputs[name] = input
	}
}

func MakeAddParam(name string, p *param.Param) Mod {
	return func(n *Node) {
		n.Params[name] = p
	}
}
