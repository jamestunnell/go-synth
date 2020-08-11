package processors

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/unit"
)

type MulK struct {
	k      float64
	kBuf   *unit.Buffer
	xBuf   *unit.Buffer
	outBuf *unit.Buffer
}

var (
	MulKPlugin = &unit.Plugin{
		BasicInfo: &unit.BasicInfo{
			Name:        "mulk",
			Description: "Multiplies a signal by a constant",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("12059b14-682c-4179-abea-aaa9b690d178"),
		},
		NewUnit: func() unit.Unit {
			return &MulK{}
		},
		GetInterface: func(srate float64) *unit.Interface {
			return &unit.Interface{
				Parameters: map[string]*unit.Parameter{
					ParamNameK: &unit.Parameter{
						Description: "multiply constant",
						Default:     1.0,
					},
				},
				NumInputs:  1,
				NumOutputs: 1,
			}
		},
		ExtraInfo: map[string]string{},
	}
)

func (mul *MulK) Initialize(
	srate float64,
	paramBuffers map[string]*unit.Buffer,
	inBuffers,
	outBuffers []*unit.Buffer) error {
	kBuf, err := unit.FindNamedBuffer(paramBuffers, ParamNameK)
	if err != nil {
		return err
	}

	if len(inBuffers) < 1 {
		return errors.New("missing input")
	}

	if len(outBuffers) < 1 {
		return errors.New("missing output")
	}

	mul.kBuf = kBuf
	mul.xBuf = inBuffers[0]
	mul.outBuf = outBuffers[0]

	return nil
}

func (mul *MulK) Configure() {
	mul.k = mul.kBuf.Values[0]
}

func (mul *MulK) Sample() {
	for i := 0; i < mul.outBuf.Length; i++ {
		mul.outBuf.Values[i] = mul.xBuf.Values[i] * mul.k
	}
}
