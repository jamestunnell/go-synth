package add

import "github.com/jamestunnell/go-synth/node"

type Add struct {
	in1Buf, in2Buf *node.Buffer
}

func init() {
	node.WorkingRegistry().RegisterCore(New())
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := node.Map{"In1": in1, "In2": in2}
	return node.NewNode(New(), inputs, node.Map{})
}

func New() *Add {
	return &Add{}
}

func (a *Add) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{"In1", "In2"},
		ControlDefaults: map[string]float64{},
	}
}

func (a *Add) Initialize(srate float64, inputs, controls node.Map) {
	a.in1Buf = inputs["In1"].Output()
	a.in2Buf = inputs["In2"].Output()
}

func (a *Add) Configure() {
}

func (a *Add) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = a.in1Buf.Values[i] + a.in2Buf.Values[i]
	}
}
