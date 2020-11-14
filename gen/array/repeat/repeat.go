package repeat

import "github.com/jamestunnell/go-synth/node"

type Repeat struct {
	Values []float64 `json:"values"`

	idx int
}

func NewNode(vals []float64) *node.Node {
	return node.NewNode(New(vals), node.Map{}, node.Map{})
}

func New(vals []float64) *Repeat {
	return &Repeat{
		Values: vals,
		idx:    0,
	}
}

func init() {
	node.WorkingRegistry().RegisterCore(&Repeat{})
}

func (r *Repeat) Interface() *node.Interface {
	return node.NewInterface()
}

func (r *Repeat) Initialize(srate float64, inputs, controls node.Map) {
}

func (r *Repeat) Configure() {
}

func (r *Repeat) Run(out *node.Buffer) {
	n := len(r.Values)

	if n == 0 {
		for i := 0; i < out.Length; i++ {
			out.Values[i] = 0.0
		}
	}

	for i := 0; i < out.Length; i++ {
		out.Values[i] = r.Values[r.idx%n]
		r.idx++
	}
}
