package array

import (
	"fmt"

	"github.com/jamestunnell/go-synth/util/param"

	"github.com/jamestunnell/go-synth/node"
)

const (
	// ParamNameRepeat is the name of the repeat param
	ParamNameRepeat = "Repeat"
	// ParamNameValues is the name of the values param
	ParamNameValues = "Values"
)

// Array will output the values once (with zeros thereafter) or repeatedly.
type Array struct {
	values []float64
	repeat bool
	idx    int
}

// NewOneshot makes an array whose values will be outputted once (with zeros
// thereafter).
func NewOneshot(vals []float64) *node.Node {
	return NewArray(vals, false)
}

// NewRepeat makes an array whose values will be outputted repeatedly.
func NewRepeat(vals []float64) *node.Node {
	return NewArray(vals, true)
}

// NewArray makes an array node
func NewArray(vals []float64, repeat bool) *node.Node {
	return node.New(&Array{},
		node.MakeAddParam(ParamNameValues, param.NewFloats(vals)),
		node.MakeAddParam(ParamNameRepeat, param.NewBool(repeat)))
}

// Interface provides the node interface.
func (a *Array) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.ParamTypes[ParamNameRepeat] = param.Bool
	ifc.ParamTypes[ParamNameValues] = param.Floats

	return ifc
}

// Initialize initializes the node.
// Returns a non-nil error if values are empty.
func (a *Array) Initialize(args *node.InitArgs) error {
	a.repeat = args.Params[ParamNameRepeat].Value().(bool)
	a.values = args.Params[ParamNameValues].Value().([]float64)

	if len(a.values) == 0 {
		return fmt.Errorf("%s param is empty", ParamNameValues)
	}

	a.idx = 0

	return nil
}

// Configure does nothing.
func (a *Array) Configure() {
}

// Run outputs array values or zeros.
func (a *Array) Run(out *node.Buffer) {
	if a.repeat {
		n := len(a.values)
		totalCopied := 0

		for totalCopied < out.Length {
			// This function copies the minimum of len(dst) and len(src) so we
			// should be safe to try copying as much as possible each time
			nCopied := copy(out.Values[totalCopied:], a.values[a.idx:])

			totalCopied += nCopied
			a.idx += nCopied

			if a.idx > (n - 1) {
				a.idx = 0
			}
		}
	} else {
		var nCopied int
		if a.idx < len(a.values) {
			// This function copies the minimum of len(dst) and len(src) so we
			// should be safe to try copying as much as possible each time
			nCopied = copy(out.Values, a.values[a.idx:])

			a.idx += nCopied
		}

		if nCopied < len(out.Values) {
			Fill(out.Values[nCopied:], 0.0)
		}
	}
}
