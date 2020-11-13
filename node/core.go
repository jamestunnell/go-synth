package node

//go:generate mockgen -source core.go -destination nodetest/mockcore.go -package nodetest

// Core defines the functional core of a node.
type Core interface {
	Interface() *Interface
	Initialize(srate float64, inputs, controls Map)
	Configure()
	Run(*Buffer)
}
