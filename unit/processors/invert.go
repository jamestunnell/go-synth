package processors

import (
	"errors"
	"math"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/unit"
)

type Invert struct {
	inBuf  *unit.Buffer
	outBuf *unit.Buffer
}

var (
	InvertPlugin = &unit.Plugin{
		BasicInfo: &unit.BasicInfo{
			Name:        "Invert",
			Description: "Invert a signal (multiplicative inverse)",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("128c35b2-ee15-4368-b42a-c99758c742af"),
		},
		NewUnit: func() unit.Unit {
			return &Invert{}
		},
		GetInterface: func(srate float64) *unit.Interface {
			return &unit.Interface{
				NumInputs:  1,
				NumOutputs: 1,
			}
		},
		ExtraInfo: map[string]string{},
	}
)

func (inv *Invert) Initialize(
	srate float64,
	paramBuffers map[string]*unit.Buffer,
	inBuffers,
	outBuffers []*unit.Buffer) error {
	if len(inBuffers) < 1 {
		return errors.New("missing input")
	}

	if len(outBuffers) < 1 {
		return errors.New("missing output")
	}

	inv.inBuf = inBuffers[0]
	inv.outBuf = outBuffers[0]

	return nil
}

func (inv *Invert) Configure() {
}

func (inv *Invert) Sample() {
	for i := 0; i < inv.outBuf.Length; i++ {
		x := inv.inBuf.Values[i]

		var y float64
		if x == 0.0 {
			y = math.MaxFloat64
		} else {
			y = 1.0 / x
		}

		inv.outBuf.Values[i] = y
	}
}
