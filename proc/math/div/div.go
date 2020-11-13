package div

import "github.com/jamestunnell/go-synth/node"

type Div struct {
	in1Buf, in2Buf *node.Buffer
}

func init() {
	node.WorkingRegistry().RegisterCore(New())
}

func NewNode(in1, in2 *node.Node) *node.Node {
	inputs := node.Map{"In1": in1, "In2": in2}
	return node.NewNode(New(), inputs, node.Map{})
}

func New() *Div {
	return &Div{}
}

func (d *Div) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{"In1", "In2"},
		ControlDefaults: map[string]float64{},
	}
}

func (d *Div) Initialize(srate float64, inputs, controls node.Map) {
	d.in1Buf = inputs["In1"].Output()
	d.in2Buf = inputs["In2"].Output()
}

func (d *Div) Configure() {
}

func (d *Div) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = d.in1Buf.Values[i] / d.in2Buf.Values[i]
	}
}
