package node

type Node struct {
	Out        *Buffer
	Core       Core
	InputNodes []*Node
	ParamNodes []*Node
}

func (n *Node) Initialize(srate float64) {
	for _, inputNode := range n.InputNodes {
		inputNode.Initialize(srate)
	}

	paramSampleRate := srate / float64(n.Out.Length)
	for _, paramNode := range n.ParamNodes {
		paramNode.Initialize(paramSampleRate)
	}

	n.Core.Initialize(srate)
}

func (n *Node) Sample() {
	for _, inputNode := range n.InputNodes {
		inputNode.Sample()
	}

	for _, paramNode := range n.ParamNodes {
		paramNode.Sample()
	}

	n.Core.Configure()
	n.Core.Sample(n.Out)
}
