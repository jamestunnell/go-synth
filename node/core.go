package node

type Core interface {
	Initialize(srate float64, inputs, controls map[string]*Node)
	Configure()
	Run(*Buffer)
}
