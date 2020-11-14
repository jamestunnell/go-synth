package osc

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type RunOscFunc func(phase float64) float64

type Osc struct {
	freqBuf  *node.Buffer
	phaseBuf *node.Buffer

	srate       float64
	phaseIncr   float64
	phase       float64
	phaseOffset float64
}

const (
	ControlFreq  = "Freq"
	ControlPhase = "Phase"
	twoPi        = 2.0 * math.Pi
)

func NewNode(core node.Core, freq, phase *node.Node) *node.Node {
	controls := node.Map{
		ControlFreq:  freq,
		ControlPhase: phase,
	}
	return node.NewNode(core, node.Map{}, controls)
}

func (osc *Osc) Interface() *node.Interface {
	return &node.Interface{
		InputNames: []string{},
		ControlDefaults: map[string]float64{
			ControlFreq:  1.0,
			ControlPhase: 0.0,
		},
	}
}

func (osc *Osc) Initialize(srate float64, inputs, controls node.Map) {
	osc.freqBuf = controls[ControlFreq].Output()
	osc.phaseBuf = controls[ControlPhase].Output()

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

func (osc *Osc) Run(runOsc RunOscFunc, out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = runOsc(osc.phase)

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
