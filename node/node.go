package node

type Node struct {
	Core        Core
	Inputs      map[string]*Node
	Controls    map[string]*Node
	Output      *Buffer
	Initialized bool
}

func New(core Core, inputs, controls map[string]*Node) *Node {
	return &Node{
		Core:        core,
		Inputs:      inputs,
		Controls:    controls,
		Output:      nil,
		Initialized: false,
	}
}

func (n *Node) Initialize(srate float64, depth int) {
	for _, inputNode := range n.Inputs {
		inputNode.Initialize(srate, depth)
	}

	controlSampleRate := srate / float64(depth)
	for _, controlNode := range n.Controls {
		controlNode.Initialize(controlSampleRate, 1)
	}

	n.Core.Initialize(srate, n.Inputs, n.Controls)

	n.Output = NewBuffer(depth)
	n.Initialized = true
}

func (n *Node) Run() {
	if !n.Initialized {
		panic("node is not initialized")
	}

	for _, inputNode := range n.Inputs {
		inputNode.Run()
	}

	for _, controlNode := range n.Controls {
		controlNode.Run()
	}

	n.Core.Configure()
	n.Core.Run(n.Output)
}
