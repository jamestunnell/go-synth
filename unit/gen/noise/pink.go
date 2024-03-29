package noise

import (
	"fmt"
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
	*White
	*state
}

// scaling is applied to make the final output range close to [-1,1)
const finalPinkScaling = 0.10

// New makes a new pink noise block.
func NewPink() *Pink {
	return &Pink{
		White: NewWhite(),
		state: &state{},
	}
}

// Initialize initializes the node.
func (p *Pink) Initialize(srate float64, outDepth int) error {
	if err := p.White.Initialize(srate, outDepth); err != nil {
		return fmt.Errorf("failed to init white noise: %w", err)
	}

	return nil
}

// Run generates pink noise in the range [-1.0,1.0).
func (p *Pink) Run() {
	// generate the white noise
	p.White.Run()

	for i := 0; i < len(p.Out.Buffer); i++ {
		white := p.Out.Buffer[i]

		p.state.b0 = 0.99886*p.state.b0 + white*0.0555179
		p.state.b1 = 0.99332*p.state.b1 + white*0.0750759
		p.state.b2 = 0.96900*p.state.b2 + white*0.1538520
		p.state.b3 = 0.86650*p.state.b3 + white*0.3104856
		p.state.b4 = 0.55000*p.state.b4 + white*0.5329522
		p.state.b5 = -0.7616*p.state.b5 - white*0.0168980

		pink := p.state.b0 + p.state.b1 + p.state.b2 + p.state.b3 +
			p.state.b4 + p.state.b5 + p.state.b6 + white*0.5362

		p.Out.Buffer[i] = pink * finalPinkScaling // (roughly) compensate for gain

		p.state.b6 = white * 0.115926
	}
}
