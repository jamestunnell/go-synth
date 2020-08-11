package generators

import (
	"errors"
	"math"

	"github.com/jamestunnell/go-synth/pkg/unit"
	"github.com/jamestunnell/go-synth/pkg/unit/constraints"
)

type RunOscFunc func(phase float64) float64

type Oscillator struct {
	srate          float64
	phaseIncr      float64
	phase          float64
	freqBuf        *unit.Buffer
	phaseOffsetBuf *unit.Buffer
	outBuf         *unit.Buffer
	runOsc         RunOscFunc
}

const (
	ParamNameFreq        = "freq"
	ParamNamePhaseOffset = "phaseOffset"
	TwoPi                = 2.0 * math.Pi
)

func NewOscillator(f RunOscFunc) *Oscillator {
	return &Oscillator{runOsc: f}
}

func NewOscillatorPlugin(basicInfo *unit.BasicInfo, f RunOscFunc) *unit.Plugin {
	return &unit.Plugin{
		BasicInfo: basicInfo,
		NewUnit:   func() unit.Unit { return NewOscillator(f) },
		GetInterface: func(srate float64) *unit.Interface {
			return &unit.Interface{
				Parameters: []*unit.Parameter{
					&unit.Parameter{
						Name:        ParamNameFreq,
						Description: "frequency",
						Required:    true,
						Constraints: []unit.Constraint{
							constraints.NewGreater(0.0),
							constraints.NewLessEqual(srate / 2.0),
						},
					},
					&unit.Parameter{
						Name:        ParamNamePhaseOffset,
						Description: "phase offset",
						Default:     0.0,
						Constraints: []unit.Constraint{
							constraints.NewGreaterEqual(-TwoPi),
							constraints.NewLessEqual(TwoPi),
						},
					},
				},
				NumOutputs: 1,
			}
		},
		ExtraInfo: map[string]string{},
	}

}

func (osc *Oscillator) Initialize(
	srate float64,
	paramBuffers map[string]*unit.Buffer,
	inBuffers,
	outBuffers []*unit.Buffer) error {
	freqBuf, err := unit.FindNamedBuffer(paramBuffers, ParamNameFreq)
	if err != nil {
		return err
	}

	phaseOffsetBuf, err := unit.FindNamedBuffer(paramBuffers, ParamNamePhaseOffset)
	if err != nil {
		return err
	}

	if len(outBuffers) < 1 {
		return errors.New("missing output")
	}

	osc.srate = srate
	osc.freqBuf = freqBuf
	osc.phaseOffsetBuf = phaseOffsetBuf
	osc.outBuf = outBuffers[0]

	osc.Configure()

	return nil
}

func (osc *Oscillator) Configure() {
	freq := osc.freqBuf.Values[0]
	phaseOffset := osc.phaseOffsetBuf.Values[0]

	osc.phaseIncr = (freq * TwoPi) / osc.srate
	osc.phase = phaseOffset
}

func (osc *Oscillator) Sample() {
	for i := 0; i < osc.outBuf.Length; i++ {
		osc.outBuf.Values[i] = osc.runOsc(osc.phase)

		osc.phase += osc.phaseIncr
		if osc.phase > math.Pi {
			osc.phase -= TwoPi
		}
	}
}
