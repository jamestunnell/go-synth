package osc

import (
	"math"

	"github.com/jamestunnell/go-synth/gen/constant"
	"github.com/jamestunnell/go-synth/node"
)

type RunOscFunc func(phase float64) float64

type Osc struct {
	freqBuf  *node.Buffer
	phaseBuf *node.Buffer

	runOsc RunOscFunc

	srate       float64
	phaseIncr   float64
	phase       float64
	phaseOffset float64
}

const twoPi = 2.0 * math.Pi

func NewNode(freq, phase *node.Node, runOsc RunOscFunc) *node.Node {
	controls := map[string]*node.Node{
		"Freq":  freq,
		"Phase": phase,
	}
	return node.New(New(runOsc), map[string]*node.Node{}, controls)
}

func New(runOsc RunOscFunc) *Osc {
	return &Osc{runOsc: runOsc}
}

func (osc *Osc) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	if n, found := controls["Freq"]; found {
		osc.freqBuf = n.Output
	} else {
		controls["Freq"] = constant.NewNode(1.0)
	}

	if n, found := controls["Phase"]; found {
		osc.phaseBuf = n.Output
	} else {
		controls["Phase"] = constant.NewNode(0.0)
	}

	osc.srate = srate
	osc.phase = 0.0
	osc.phaseOffset = 0.0
}

func (osc *Osc) Configure() {
	freq := osc.freqBuf.Values[0]
	phaseOffset := osc.phaseBuf.Values[0]

	osc.phaseIncr = (freq * twoPi) / osc.srate

	if phaseOffset != osc.phaseOffset {
		phaseOffset = processPhaseOffset(phaseOffset)
		osc.phase += phaseOffset - osc.phaseOffset
	}
}

func (osc *Osc) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = osc.runOsc(osc.phase)

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
