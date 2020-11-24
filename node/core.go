package node

import "github.com/jamestunnell/go-synth/util/param"

//go:generate mockgen -source core.go -destination nodetest/mockcore.go -package nodetest

// InitArgs is used for initializing a node.
type InitArgs struct {
	SampleRate       float64
	Params           param.Map
	Inputs, Controls Map
}

// Core defines the functional core of a node.
type Core interface {
	Interface() *Interface
	Initialize(*InitArgs) error
	Configure()
	Run(*Buffer)
}
