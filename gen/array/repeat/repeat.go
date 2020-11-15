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

	if n > 0 {
		totalCopied := 0

		for totalCopied < out.Length {
			// This function copies the minimum of len(dst) and len(src) so we
			// should be safe to try copying as much as possible each time
			nCopied := copy(out.Values[totalCopied:], r.Values[r.idx:])

			totalCopied += nCopied
			r.idx += nCopied

			if r.idx > (n - 1) {
				r.idx = 0
			}
		}
	} else {
		for i := 0; i < out.Length; i++ {
			out.Values[i] = 0.0
		}
	}
}
