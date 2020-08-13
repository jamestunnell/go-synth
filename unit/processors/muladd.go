package processors

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/unit"
)

type MulAdd struct {
	inXBuf *unit.Buffer
	inYBuf *unit.Buffer
	inZBuf *unit.Buffer
	outBuf *unit.Buffer
}

var (
	MulAddPlugin = &unit.Plugin{
		BasicInfo: &unit.BasicInfo{
			Name:        "muladd",
			Description: "Multiply first two inputs then add result to the third input",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("975075f2-0845-4344-994a-ea05a7c94dd6"),
		},
		NewUnit: func() unit.Unit {
			return &MulAdd{}
		},
		GetInterface: func(srate float64) *unit.Interface {
			return &unit.Interface{
				NumInputs:  3,
				NumOutputs: 1,
			}
		},
		ExtraInfo: map[string]string{},
	}
)

func (muladd *MulAdd) Initialize(
	srate float64,
	paramBuffers map[string]*unit.Buffer,
	inBuffers,
	outBuffers []*unit.Buffer) error {
	if len(inBuffers) < 3 {
		return errors.New("missing input(s)")
	}

	if len(outBuffers) < 1 {
		return errors.New("missing output")
	}

	muladd.inXBuf = inBuffers[0]
	muladd.inYBuf = inBuffers[1]
	muladd.inZBuf = inBuffers[2]
	muladd.outBuf = outBuffers[0]

	return nil
}

func (muladd *MulAdd) Configure() {
}

func (muladd *MulAdd) Sample() {
	for i := 0; i < muladd.outBuf.Length; i++ {
		muladd.outBuf.Values[i] = muladd.inXBuf.Values[i]*muladd.inYBuf.Values[i] + muladd.inZBuf.Values[i]
	}
}
