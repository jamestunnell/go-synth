package array

import (
	"github.com/jamestunnell/go-synth"
)

type Repeat struct {
	Values *synth.TypedParam[[]float64]
	Out    *synth.TypedOutput[float64]

	idx    int
	outBuf []float64
}

func NewRepeat(vals ...float64) *Repeat {
	r := &Repeat{
		Values: synth.NewFloat64ArrayParam(vals),
	}

	r.Out = synth.NewFloat64Output(r)

	return r
}

// Initialize initializes the block.
// Returns a non-nil error if values are empty.
func (r *Repeat) Initialize(srate float64, outDepth int) error {
	if len(r.Values.Value) == 0 {
		return errValuesEmpty
	}

	r.Out.Initialize(outDepth)

	r.idx = 0
	r.outBuf = r.Out.Buffer().([]float64)

	return nil
}

// Configure does nothing.
func (r *Repeat) Configure() {
}

// Run outputs array values or zeros.
func (r *Repeat) Run() {
	n := len(r.Values.Value)
	totalCopied := 0

	for totalCopied < len(r.outBuf) {
		// This function copies the minimum of len(dst) and len(src) so we
		// should be safe to try copying as much as possible each time
		nCopied := copy(r.outBuf[totalCopied:], r.Values.Value[r.idx:])

		totalCopied += nCopied
		r.idx += nCopied

		if r.idx > (n - 1) {
			r.idx = 0
		}
	}
}
