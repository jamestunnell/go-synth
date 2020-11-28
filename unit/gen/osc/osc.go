package osc

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-synth/util/param"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/mod"
)

// Wave is used to select the oscillator wave type
type Wave int64

type runOscFunc func(phase float64) float64

// Osc is a simple (naive) oscillator
type Osc struct {
	freqBuf  *node.Buffer
	phaseBuf *node.Buffer

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
func NewSine(freq, phase *node.Node, mods ...node.Mod) *node.Node {
	return New(Sine, freq, phase, mods...)
}

// NewSquare makes a square wave oscillator node.
func NewSquare(freq, phase *node.Node, mods ...node.Mod) *node.Node {
	return New(Square, freq, phase, mods...)
}

// NewSawtooth makes a sawtooth wave oscillator node.
func NewSawtooth(freq, phase *node.Node, mods ...node.Mod) *node.Node {
	return New(Sawtooth, freq, phase, mods...)
}

// NewTriangle makes a triangle wave oscillator node.
func NewTriangle(freq, phase *node.Node, mods ...node.Mod) *node.Node {
	return New(Triangle, freq, phase, mods...)
}

// New makes a new Osc node.
func New(wave Wave, freq, phase *node.Node, mods ...node.Mod) *node.Node {
	mods = append(mods, mod.Control(ControlFreq, freq))
	mods = append(mods, mod.Control(ControlPhase, phase))
	mods = append(mods, mod.Param(ParamWave, param.NewInt(int64(wave))))

	return node.New(&Osc{}, mods...)
}

// Interface provides the node interface.
func (osc *Osc) Interface() *node.Interface {
	return &node.Interface{
		InputNames: []string{},
		ControlDefaults: map[string]float64{
			ControlFreq:  1.0,
			ControlPhase: 0.0,
		},
		ParamTypes: map[string]param.Type{
			ParamWave: param.Int,
		},
	}
}

// Initialize initializes the node.
// Returns a non-nil error if the wave type is unexpected.
func (osc *Osc) Initialize(args *node.InitArgs) error {
	osc.freqBuf = args.Controls[ControlFreq].Output()
	osc.phaseBuf = args.Controls[ControlPhase].Output()

	wave := Wave(args.Params[ParamWave].Value().(int64))
	switch wave {
	case Sine:
		osc.runOsc = sineWave
	case Square:
		osc.runOsc = squareWave
	case Sawtooth:
		osc.runOsc = sawtoothWave
	case Triangle:
		osc.runOsc = triangleWave
	default:
		return fmt.Errorf("unknown wave type %d", wave)
	}

	osc.srate = args.SampleRate
	osc.phase = 0.0
	osc.phaseOffset = 0.0

	return nil
}

// Configure configures the node using latest output from the
// Freq and Phase controls.
func (osc *Osc) Configure() {
	freq := osc.freqBuf.Values[0]
	phaseOffset := osc.phaseBuf.Values[0]

	osc.phaseIncr = (freq * twoPi) / osc.srate

	if phaseOffset != osc.phaseOffset {
		phaseOffset = processPhaseOffset(phaseOffset)
		osc.phase += phaseOffset - osc.phaseOffset
	}
}

// Run runs the oscillator wave function and places results in the given buffer.
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
