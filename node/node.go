package node

type Node interface {
	Buffer() *Buffer
	Controls() map[string]Node
	Inputs() map[string]Node
	Initialize(srate float64, depth int)
	Configure()
	Run()
}

func Initialize(n Node, srate float64, depth int) {
	for _, inputNode := range n.Inputs() {
		Initialize(inputNode, srate, depth)
	}

	controlSampleRate := srate / float64(depth)
	for _, controlNode := range n.Controls() {
		Initialize(controlNode, controlSampleRate, 1)
	}

	n.Initialize(srate, depth)
}

func Run(n Node) {
	for _, inputNode := range n.Inputs() {
		Run(inputNode)
	}

	for _, controlNode := range n.Controls() {
		Run(controlNode)
	}

	n.Configure()
	n.Run()
}
