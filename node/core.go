package node

//go:generate mockgen -source core.go -destination nodetest/mockcore.go -package nodetest

type InitArgs struct {
	SampleRate       float64
	Params           ParamMap
	Inputs, Controls Map
}

// Core defines the functional core of a node.
type Core interface {
	Interface() *Interface
	Initialize(*InitArgs) error
	Configure()
	Run(*Buffer)
}
