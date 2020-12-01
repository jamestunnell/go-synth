package brown

import (
	"time"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen/noise/white"
)

type state struct {
	smooth float64
}

// Brown produces brown noise by running white noise through a lowpass filter.
// Adapted from https://github.com/alessandrocuda/noise_generator
// Output is from -1 to 1.
type Brown struct {
	*white.White
	*state
}

const lpfBeta = 0.025

// New makes a new White node
func New(moreMods ...node.Mod) *node.Node {
	seed := time.Now().UTC().UnixNano()

	return NewFromSeed(seed, moreMods...)
}

// NewFromSeed makes a new White node that uses the given seed.
func NewFromSeed(seed int64, moreMods ...node.Mod) *node.Node {
	mods := append(white.ParamMods(seed), moreMods...)

	return node.New(&Brown{}, mods...)
}

// Initialize initializes the node.
func (p *Brown) Initialize(args *node.InitArgs) error {
	p.White = &white.White{}
	p.state = &state{}

	return p.White.Initialize(args)
}

// Run runs the brown noise generation process.
func (p *Brown) Run(out *node.Buffer) {
	// generate the white noise
	p.White.Run(out)

	for i := 0; i < out.Length; i++ {
		white := out.Values[i]
		smooth := p.state.smooth

		p.state.smooth = smooth - (lpfBeta * (smooth - white)) // RC Filter
	}
}
