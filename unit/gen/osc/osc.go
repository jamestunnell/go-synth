package osc

import (
	"math"

	"github.com/jamestunnell/go-synth"
)

// Wave is used to select the oscillator wave type
type Wave int64

type runOscFunc func(phase float64) float64

// Osc is a simple (naive) oscillator
type Osc struct {
	Freq  *synth.TypedControl[float64]
	Phase *synth.TypedControl[float64]
	Out   *synth.TypedOutput[float64]

	outBuf   []float64
	freqBuf  []float64
	phaseBuf []float64

	runOsc runOscFunc

	srate       float64
	phaseIncr   float64
	phase       float64
	phaseOffset float64
}

const (
	// Sine selects a sine wave
	Sine Wave = iota
	// Square selects a square wave
	Square
	// Sawtooth selects a sawtooth wave
	Sawtooth
	// Triangle selects a triangle wave
	Triangle

	// ControlFreq is the name of the Freq control
	ControlFreq = "Freq"
	// ControlPhase is the name of the Phase control
	ControlPhase = "Phase"
	// ParamWave is the name of the Wave param
	ParamWave = "Wave"

	twoPi = 2.0 * math.Pi
)

// NewSine makes a sine wave oscillator node.
func NewSine() *Osc {
	return New(sineWave)
}

// NewSquare makes a square wave oscillator node.
func NewSquare() *Osc {
	return New(squareWave)
}

// NewSawtooth makes a sawtooth wave oscillator node.
func NewSawtooth() *Osc {
	return New(sawtoothWave)
}

// NewTriangle makes a triangle wave oscillator node.
func NewTriangle() *Osc {
	return New(triangleWave)
}

// New makes a new Osc node.
func New(runOsc runOscFunc) *Osc {
	osc := &Osc{
		runOsc: runOsc,
		Freq:   synth.NewTypedControl[float64](1.0),
		Phase:  synth.NewTypedControl[float64](0.0),
	}

	osc.Out = synth.NewTypedOutput[float64](osc)

	return osc
}

// Initialize initializes the node.
// Returns a non-nil error if the wave type is unexpected.
func (osc *Osc) Initialize(srate float64, outDepth int) error {
	osc.Out.Initialize(outDepth)

	osc.freqBuf = osc.Freq.Output.Buffer().([]float64)
	osc.phaseBuf = osc.Phase.Output.Buffer().([]float64)
	osc.outBuf = osc.Out.Buffer().([]float64)
	osc.srate = srate
	osc.phase = 0.0
	osc.phaseOffset = 0.0

	return nil
}

// Configure configures the node using latest output from the
// Freq and Phase controls.
func (osc *Osc) Configure() {
	osc.phaseIncr = (osc.freqBuf[0] * twoPi) / osc.srate

	phaseOffset := osc.phaseBuf[0]
	if phaseOffset != osc.phaseOffset {
		phaseOffset = processPhaseOffset(phaseOffset)
		osc.phase += phaseOffset - osc.phaseOffset
	}
}

// Run runs the oscillator wave function and places results in the given buffer.
func (osc *Osc) Run() {
	for i := 0; i < len(osc.outBuf); i++ {
		osc.outBuf[i] = osc.runOsc(osc.phase)

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
