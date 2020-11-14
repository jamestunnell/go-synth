package oneshot

import "github.com/jamestunnell/go-synth/node"

type Oneshot struct {
	Values []float64 `json:"values"`

	idx int
}

func init() {
	node.WorkingRegistry().RegisterCore(&Oneshot{})
}

func NewNode(vals []float64) *node.Node {
	return node.NewNode(New(vals), node.Map{}, node.Map{})
}

func New(vals []float64) *Oneshot {
	return &Oneshot{
		Values: vals,
		idx:    0,
	}
}

func (o *Oneshot) Interface() *node.Interface {
	return node.NewInterface()
}

func (o *Oneshot) Initialize(srate float64, inputs, controls node.Map) {
}

func (o *Oneshot) Configure() {
}

func (o *Oneshot) Run(out *node.Buffer) {
	n := len(o.Values)

	if n == 0 {
		for i := 0; i < out.Length; i++ {
			out.Values[i] = 0.0
		}
	}

	for i := 0; i < out.Length; i++ {
		var outVal float64
		if o.idx < n {
			outVal = o.Values[o.idx]
			o.idx++
		}
		out.Values[i] = outVal
	}
}
