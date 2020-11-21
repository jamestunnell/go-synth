package oneshot

import (
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen/array"
)

type Oneshot struct {
	Values []float64 `json:"values"`

	idx int
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
	var nCopied int
	if o.idx < len(o.Values) {
		// This function copies the minimum of len(dst) and len(src) so we
		// should be safe to try copying as much as possible each time
		nCopied = copy(out.Values, o.Values[o.idx:])

		o.idx += nCopied
	}

	if nCopied < len(out.Values) {
		array.Fill(out.Values[nCopied:], 0.0)
	}
}
