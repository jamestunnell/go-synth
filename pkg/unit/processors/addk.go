package processors

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/pkg/unit"
)

type AddK struct {
	k      float64
	kBuf   *unit.Buffer
	xBuf   *unit.Buffer
	outBuf *unit.Buffer
}

const ParamNameK = "k"

var (
	AddKPlugin = &unit.Plugin{
		BasicInfo: &unit.BasicInfo{
			Name:        "addk",
			Description: "Add a constant to a signal",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("7bbdf9f2-b323-46be-bd0c-39068e06a94a"),
		},
		NewUnit: func() unit.Unit {
			return &AddK{}
		},
		GetInterface: func(srate float64) *unit.Interface {
			return &unit.Interface{
				Parameters: map[string]*unit.Parameter{
					ParamNameK: &unit.Parameter{
						Description: "add constant",
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

func (add *AddK) Initialize(
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

	add.kBuf = kBuf
	add.xBuf = inBuffers[0]
	add.outBuf = outBuffers[0]

	return nil
}

func (add *AddK) Configure() {
	add.k = add.kBuf.Values[0]
}

func (add *AddK) Sample() {
	for i := 0; i < add.outBuf.Length; i++ {
		add.outBuf.Values[i] = add.xBuf.Values[i] + add.k
	}
}
