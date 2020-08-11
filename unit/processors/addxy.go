package processors

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/unit"
)

type AddXY struct {
	inXBuf *unit.Buffer
	inYBuf *unit.Buffer
	outBuf *unit.Buffer
}

var (
	AddXYPlugin = &unit.Plugin{
		BasicInfo: &unit.BasicInfo{
			Name:        "addxy",
			Description: "Adds two signals",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("1e62b892-5c86-4022-87e9-6560630890e4"),
		},
		NewUnit: func() unit.Unit {
			return &AddXY{}
		},
		GetInterface: func(srate float64) *unit.Interface {
			return &unit.Interface{
				Parameters: map[string]*unit.Parameter{},
				NumInputs:  2,
				NumOutputs: 1,
			}
		},
		ExtraInfo: map[string]string{},
	}
)

func (add *AddXY) Initialize(
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

	add.inXBuf = inBuffers[0]
	add.inYBuf = inBuffers[1]
	add.outBuf = outBuffers[0]

	return nil
}

func (add *AddXY) Configure() {
}

func (add *AddXY) Sample() {
	for i := 0; i < add.outBuf.Length; i++ {
		add.outBuf.Values[i] = add.inXBuf.Values[i] + add.inYBuf.Values[i]
	}
}
