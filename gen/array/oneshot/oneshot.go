package oneshot

import "github.com/jamestunnell/go-synth/node"

type Oneshot struct {
	vals []float64

	idx     int
	numVals int
}

func NewNode(vals []float64) *node.Node {
	return node.New(New(vals), map[string]*node.Node{}, map[string]*node.Node{})
}

func New(vals []float64) *Oneshot {
	if len(vals) == 0 {
		panic("Oneshot has no values")
	}

	o := &Oneshot{
		vals:    vals,
		numVals: len(vals),
		idx:     0,
	}

	return o
}

func (o *Oneshot) Initialize(srate float64, inputs, controls map[string]*node.Node) {
}

func (o *Oneshot) Configure() {
}

func (o *Oneshot) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		var outVal float64
		if o.idx < o.numVals {
			outVal = o.vals[o.idx]
			o.idx++
		}
		out.Values[i] = outVal
	}
}
