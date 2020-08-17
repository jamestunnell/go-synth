package generators

import (
	"errors"
	"math"

	"github.com/jamestunnell/go-synth/unit"
)

type RunOscFunc func(phase float64) float64

type Oscillator struct {
	srate          float64
	ampl           float64
	phaseIncr      float64
	phase          float64
	freqBuf        *unit.Buffer
	phaseOffsetBuf *unit.Buffer
	amplBuf        *unit.Buffer
	outBuf         *unit.Buffer
	runOsc         RunOscFunc
}

const (
	ParamNameFreq      = "Freq"
	ParamNamePhase     = "Phase"
	ParamNameAmplitude = "Ampl"
	TwoPi              = 2.0 * math.Pi
)

func NewOscillator(f RunOscFunc) *Oscillator {
	return &Oscillator{runOsc: f}
}

func NewOscillatorPlugin(basicInfo *unit.BasicInfo, f RunOscFunc) *unit.Plugin {
	return &unit.Plugin{
		BasicInfo: basicInfo,
		Interface: &unit.Interface{
			Parameters: map[string]*unit.ParamInfo{
				ParamNameFreq: &unit.ParamInfo{
					Description: "frequency",
					Required:    true,
					NVConstraints: []unit.NVConstraintInfo{
						unit.Positive.Info(),
						unit.NyquistLimited.Info(),
					},
				},
				ParamNamePhase: &unit.ParamInfo{
					Description: "phase offset",
					Default:     0.0,
				},
				ParamNameAmplitude: &unit.ParamInfo{
					Description: "amplitude",
					Default:     1.0,
				},
			},
			NumOutputs: 1,
		},
		NewUnit:   func() unit.Unit { return NewOscillator(f) },
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

	phaseOffsetBuf, err := unit.FindNamedBuffer(paramBuffers, ParamNamePhase)
	if err != nil {
		return err
	}

	amplBuf, err := unit.FindNamedBuffer(paramBuffers, ParamNameAmplitude)
	if err != nil {
		return err
	}

	if len(outBuffers) < 1 {
		return errors.New("missing output")
	}

	osc.srate = srate
	osc.freqBuf = freqBuf
	osc.phaseOffsetBuf = phaseOffsetBuf
	osc.amplBuf = amplBuf
	osc.outBuf = outBuffers[0]

	return nil
}

func (osc *Oscillator) Configure() {
	freq := osc.freqBuf.Values[0]
	phaseOffset := osc.phaseOffsetBuf.Values[0]

	processPhaseOffset := func(x float64) float64 {
		var y float64

		// This will put the phase in range [-TwoPi,TwoPi] if needed
		if x > TwoPi && x < -TwoPi {
			y = math.Mod(x, TwoPi)
		} else {
			y = x
		}

		// Move phase from [Pi,TwoPi] -> [-Pi,0]
		if y > math.Pi {
			return y - TwoPi
		}

		// Move phase from [-TwoPi,-Pi] -> [0,Pi]
		if y < -math.Pi {
			return y + TwoPi
		}

		return y
	}

	osc.phaseIncr = (freq * TwoPi) / osc.srate
	osc.phase = processPhaseOffset(phaseOffset)
	osc.ampl = osc.amplBuf.Values[0]
}

func (osc *Oscillator) Sample() {
	for i := 0; i < osc.outBuf.Length; i++ {
		osc.outBuf.Values[i] = osc.ampl * osc.runOsc(osc.phase)

		osc.phase += osc.phaseIncr
		if osc.phase > math.Pi {
			osc.phase -= TwoPi
		}
	}
}
