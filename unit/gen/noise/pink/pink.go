package pink

import (
	"time"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen/noise/white"
)

type state struct {
	b0, b1, b2, b3, b4, b5, b6 float64
}

// Pink produces pink (1/f) noise by running white noise through a lowpass filter.
// Pink noise filtering based on Paul Kellet's "Filter to make pink noise from white"
// posted on musicdsp.org.
// See https://www.musicdsp.org/en/latest/Filters/76-pink-noise-filter.html
// Output is from -1 to 1.
type Pink struct {
	*white.White
	*state
}

// scaling is applied to keep the final output range to be within [-1,1)
const finalScaling = 0.099

// New makes a new White node
func New(moreMods ...node.Mod) *node.Node {
	seed := time.Now().UTC().UnixNano()

	return NewFromSeed(seed, moreMods...)
}

// NewFromSeed makes a new White node that uses the given seed.
func NewFromSeed(seed int64, moreMods ...node.Mod) *node.Node {
	mods := append(white.ParamMods(seed), moreMods...)

	return node.New(&Pink{}, mods...)
}

// Initialize initializes the node.
func (p *Pink) Initialize(args *node.InitArgs) error {
	p.White = &white.White{}
	p.state = &state{}

	return p.White.Initialize(args)
}

// Run runs the pink noise generation process.
func (p *Pink) Run(out *node.Buffer) {
	// generate the white noise
	p.White.Run(out)

	for i := 0; i < out.Length; i++ {
		white := out.Values[i]

		p.state.b0 = 0.99886*p.state.b0 + white*0.0555179
		p.state.b1 = 0.99332*p.state.b1 + white*0.0750759
		p.state.b2 = 0.96900*p.state.b2 + white*0.1538520
		p.state.b3 = 0.86650*p.state.b3 + white*0.3104856
		p.state.b4 = 0.55000*p.state.b4 + white*0.5329522
		p.state.b5 = -0.7616*p.state.b5 - white*0.0168980

		pink := p.state.b0 + p.state.b1 + p.state.b2 + p.state.b3 +
			p.state.b4 + p.state.b5 + p.state.b6 + white*0.5362

		out.Values[i] = pink * finalScaling // (roughly) compensate for gain

		p.state.b6 = white * 0.115926
	}
}
