package array

import (
	"github.com/jamestunnell/go-synth"
)

type Repeat struct {
	Values *synth.TypedParam[[]float64]
	Out    *synth.TypedOutput[float64]

	idx int
}

func NewRepeatNoVals() *Repeat {
	return &Repeat{
		Values: synth.NewTypedParam([]float64{}),
		Out:    synth.NewFloat64Output(),
	}
}

func NewRepeat(vals ...float64) *Repeat {
	return &Repeat{
		Values: synth.NewTypedParam(vals),
		Out:    synth.NewFloat64Output(),
	}
}

// Initialize initializes the block.
// Returns a non-nil error if values are empty.
func (r *Repeat) Initialize(srate float64, outDepth int) error {
	if len(r.Values.Value) == 0 {
		return errValuesEmpty
	}

	r.Out.Initialize(outDepth)

	r.idx = 0

	return nil
}

// Configure does nothing.
func (r *Repeat) Configure() {
}

// Run outputs array values or zeros.
func (r *Repeat) Run() {
	n := len(r.Values.Value)
	totalCopied := 0

	for totalCopied < len(r.Out.Buffer) {
		// This function copies the minimum of len(dst) and len(src) so we
		// should be safe to try copying as much as possible each time
		nCopied := copy(r.Out.Buffer[totalCopied:], r.Values.Value[r.idx:])

		totalCopied += nCopied
		r.idx += nCopied

		if r.idx > (n - 1) {
			r.idx = 0
		}
	}
}
