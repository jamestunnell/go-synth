package node

type Core interface {
	GetInterface() *Interface
	Initialize(srate float64)
	Configure()
	Sample(out *Buffer)
}
