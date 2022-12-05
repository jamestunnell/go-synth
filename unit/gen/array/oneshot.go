package array

import (
	"errors"

	"github.com/jamestunnell/go-synth"
)

type Oneshot struct {
	Values *synth.TypedParam[[]float64]
	Out    *synth.TypedOutput[float64]

	idx int
}

var errValuesEmpty = errors.New("values param is empty")

func NewOneshot(vals ...float64) *Oneshot {
	o := &Oneshot{
		Values: synth.NewFloat64ArrayParam(vals),
	}

	o.Out = synth.NewFloat64Output(o)

	return o
}

// Initialize initializes the block.
// Returns a non-nil error if values are empty.
func (o *Oneshot) Initialize(srate float64, outDepth int) error {
	if len(o.Values.Value) == 0 {
		return errValuesEmpty
	}

	o.Out.Initialize(outDepth)

	o.idx = 0

	return nil
}

// Configure does nothing.
func (o *Oneshot) Configure() {
}

// Run outputs array values or zeros.
func (o *Oneshot) Run() {
	var nCopied int
	if o.idx < len(o.Values.Value) {
		// This function copies the minimum of len(dst) and len(src) so we
		// should be safe to try copying as much as possible each time
		nCopied = copy(o.Out.BufferValues, o.Values.Value[o.idx:])

		o.idx += nCopied
	}

	if nCopied < len(o.Out.BufferValues) {
		Fill(o.Out.BufferValues[nCopied:], 0.0)
	}
}
