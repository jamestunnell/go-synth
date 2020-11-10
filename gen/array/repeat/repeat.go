package repeat

import "github.com/jamestunnell/go-synth/node"

type Repeat struct {
	vals []float64

	idx     int
	numVals int
}

func NewNode(vals []float64) *node.Node {
	return node.New(New(vals), map[string]*node.Node{}, map[string]*node.Node{})
}

func New(vals []float64) *Repeat {
	if len(vals) == 0 {
		panic("Repeat has no values")
	}

	o := &Repeat{
		vals:    vals,
		numVals: len(vals),
		idx:     0,
	}

	return o
}

func (r *Repeat) Initialize(srate float64, inputs, controls map[string]*node.Node) {
}

func (r *Repeat) Configure() {
}

func (r *Repeat) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = r.vals[r.idx%r.numVals]
		r.idx++
	}
}
