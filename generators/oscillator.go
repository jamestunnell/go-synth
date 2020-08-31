package generators

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type OscType int

const (
	Sine OscType = iota
	Triangle
	Square
	Sawtooth
)

type runOscFunc func(phase float64) float64

type Osc struct {
	Freq   interface{}
	Phase  interface{}
	Ampl   interface{}
	Offset interface{}
	Type   OscType

	freqBuf   *node.Buffer
	phaseBuf  *node.Buffer
	amplBuf   *node.Buffer
	offsetBuf *node.Buffer

	srate       float64
	ampl        float64
	phaseIncr   float64
	phase       float64
	phaseOffset float64
	offset      float64
	runOsc      runOscFunc
}

const (
	OneOverPi = 1.0 / math.Pi
	TwoPi     = 2.0 * math.Pi

	// KSineB is used to calculate an approximation of a sine wave
	KSineB = 4.0 / math.Pi
	// KSineC is used to calculate an approximation of a sine wave
	KSineC = -4.0 / (math.Pi * math.Pi)
	// KSineP is used to calculate an approximation of a sine wave
	KSineP = 0.225
	// KTriangle is used to calculate a triangle wave
	KTriangle = 2.0 / math.Pi
)

func (osc *Osc) GetInterface() *node.Interface {
	return &node.Interface{
		Parameters: map[string]*node.ParamInfo{
			"Ampl":   &node.ParamInfo{Required: false, Default: 1.0},
			"Phase":  &node.ParamInfo{Required: false, Default: 0.0},
			"Offset": &node.ParamInfo{Required: false, Default: 0.0},
			"Freq":   &node.ParamInfo{Required: true},
		},
		Inputs: []string{},
	}
}

func (osc *Osc) Initialize(srate float64) {
	osc.srate = srate
	osc.phase = 0.0
	osc.phaseOffset = 0.0

	osc.freqBuf = osc.Freq.(*node.Node).Out
	osc.amplBuf = osc.Ampl.(*node.Node).Out
	osc.phaseBuf = osc.Phase.(*node.Node).Out
	osc.offsetBuf = osc.Offset.(*node.Node).Out

	switch osc.Type {
	case Sine:
		osc.runOsc = SineWave
	case Triangle:
		osc.runOsc = TriangleWave
	case Square:
		osc.runOsc = SquareWave
	case Sawtooth:
		osc.runOsc = SawtoothWave
	default:
		panic(fmt.Sprintf("invalid osc type", osc.Type))
	}
}

func (osc *Osc) Configure() {
	osc.ampl = osc.amplBuf.Values[0]
	osc.offset = osc.offsetBuf.Values[0]

	freq := osc.freqBuf.Values[0]
	osc.phaseIncr = (freq * TwoPi) / osc.srate

	phaseOffset := osc.phaseBuf.Values[0]

	if phaseOffset != osc.phaseOffset {
		phaseOffset = processPhaseOffset(phaseOffset)
		osc.phase += phaseOffset - osc.phaseOffset
	}
}

func (osc *Osc) Sample(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = osc.ampl*osc.runOsc(osc.phase) + osc.offset

		osc.phase += osc.phaseIncr
		if osc.phase > math.Pi {
			osc.phase -= TwoPi
		}
	}
}

func SineWave(phase float64) float64 {
	y := KSineB*phase + KSineC*phase*math.Abs(phase)
	// for extra precision
	return KSineP*(y*math.Abs(y)-y) + y // Q * y + P * y * y.abs
}

func SawtoothWave(phase float64) float64 {
	return phase * OneOverPi
}

func TriangleWave(phase float64) float64 {
	return math.Abs(KTriangle*phase) - 1.0
}

func SquareWave(phase float64) float64 {
	var y float64
	if phase >= 0.0 {
		y = 1.0
	} else {
		y = -1.0
	}

	return y
}

func processPhaseOffset(x float64) float64 {
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
