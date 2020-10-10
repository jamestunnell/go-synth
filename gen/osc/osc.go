package osc

import (
	"math"

	"github.com/jamestunnell/go-synth/gen/constant"
	"github.com/jamestunnell/go-synth/node"
)

type RunOscFunc func(phase float64) float64

type Params struct {
	Freq   node.Node
	Phase  node.Node
	Ampl   node.Node
	Offset node.Node
}

type Osc struct {
	params *Params
	runOsc RunOscFunc

	outBuf      *node.Buffer
	srate       float64
	ampl        float64
	phaseIncr   float64
	phase       float64
	phaseOffset float64
	offset      float64
}

const twoPi = 2.0 * math.Pi

func New(params *Params, runOsc RunOscFunc) *Osc {
	if params.Freq == nil {
		params.Freq = constant.New(1.0)
	}

	if params.Ampl == nil {
		params.Ampl = constant.New(1.0)
	}

	if params.Offset == nil {
		params.Offset = constant.New(0.0)
	}

	if params.Phase == nil {
		params.Phase = constant.New(0.0)
	}

	return &Osc{params: params, runOsc: runOsc}
}

func (osc *Osc) Buffer() *node.Buffer {
	return osc.outBuf
}

func (osc *Osc) Controls() map[string]node.Node {
	return map[string]node.Node{
		"Freq":   osc.params.Freq,
		"Ampl":   osc.params.Ampl,
		"Offset": osc.params.Offset,
		"Phase":  osc.params.Phase,
	}
}

func (osc *Osc) Inputs() map[string]node.Node {
	return map[string]node.Node{}
}

func (osc *Osc) Initialize(srate float64, depth int) {
	osc.outBuf = node.NewBuffer(depth)
	osc.srate = srate
	osc.phase = 0.0
	osc.phaseOffset = 0.0
}

func (osc *Osc) Configure() {
	freq := osc.params.Freq.Buffer().Values[0]
	phaseOffset := osc.params.Phase.Buffer().Values[0]

	osc.ampl = osc.params.Ampl.Buffer().Values[0]
	osc.offset = osc.params.Offset.Buffer().Values[0]
	osc.phaseIncr = (freq * twoPi) / osc.srate

	if phaseOffset != osc.phaseOffset {
		phaseOffset = processPhaseOffset(phaseOffset)
		osc.phase += phaseOffset - osc.phaseOffset
	}
}

func (osc *Osc) Run() {
	for i := 0; i < osc.outBuf.Length; i++ {
		osc.outBuf.Values[i] = osc.ampl*osc.runOsc(osc.phase) + osc.offset

		osc.phase += osc.phaseIncr
		for osc.phase > math.Pi {
			osc.phase -= twoPi
		}
	}
}

func processPhaseOffset(x float64) float64 {
	var y float64

	// This will put the phase in range [-twoPi,twoPi] if needed
	if x > twoPi && x < -twoPi {
		y = math.Mod(x, twoPi)
	} else {
		y = x
	}

	// Move phase from [Pi,twoPi] -> [-Pi,0]
	if y > math.Pi {
		return y - twoPi
	}

	// Move phase from [-twoPi,-Pi] -> [0,Pi]
	if y < -math.Pi {
		return y + twoPi
	}

	return y
}
