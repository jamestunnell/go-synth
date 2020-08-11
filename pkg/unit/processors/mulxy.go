package processors

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/pkg/unit"
)

type MulXY struct {
	inXBuf *unit.Buffer
	inYBuf *unit.Buffer
	outBuf *unit.Buffer
}

var (
	MulXYPlugin = &unit.Plugin{
		BasicInfo: &unit.BasicInfo{
			Name:        "mulxy",
			Description: "Multiplies two signals",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("358da2cf-8168-44e8-a88b-048e0a27e76f"),
		},
		NewUnit: func() unit.Unit {
			return &MulXY{}
		},
		GetInterface: func(srate float64) *unit.Interface {
			return &unit.Interface{
				Parameters: []*unit.Parameter{},
				NumInputs:  2,
				NumOutputs: 1,
			}
		},
		ExtraInfo: map[string]string{},
	}
)

func (mul *MulXY) Initialize(
	srate float64,
	paramBuffers map[string]*unit.Buffer,
	inBuffers,
	outBuffers []*unit.Buffer) error {
	if len(inBuffers) < 2 {
		return errors.New("missing input(s)")
	}

	if len(outBuffers) < 1 {
		return errors.New("missing output")
	}

	mul.inXBuf = inBuffers[0]
	mul.inYBuf = inBuffers[1]
	mul.outBuf = outBuffers[0]

	return nil
}

func (mul *MulXY) Configure() {
}

func (mul *MulXY) Sample() {
	for i := 0; i < mul.outBuf.Length; i++ {
		mul.outBuf.Values[i] = mul.inXBuf.Values[i] * mul.inYBuf.Values[i]
	}
}
