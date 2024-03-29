package noise

import (
	"fmt"
)

// Brown produces brown noise by running white noise through a lowpass filter.
// Adapted from https://github.com/alessandrocuda/noise_generator
// Output is from -1 to 1.
type Brown struct {
	*White
	smooth float64
}

// scaling is applied to make the final output range close to [-1,1)
const finalScaling = 3.0

const lpfBeta = 0.025

// New makes a new brown noise block.
func NewBrown() *Brown {
	return &Brown{
		White:  NewWhite(),
		smooth: 0,
	}
}

// Initialize initializes the node.
func (b *Brown) Initialize(srate float64, outDepth int) error {
	if err := b.White.Initialize(srate, outDepth); err != nil {
		return fmt.Errorf("failed to init white noise: %w", err)
	}

	b.smooth = 0

	return nil
}

// Run runs the brown noise generation process.
func (b *Brown) Run() {
	// generate the white noise
	b.White.Run()

	for i := 0; i < len(b.Out.Buffer); i++ {
		white := b.Out.Buffer[i]

		b.smooth = b.smooth - (lpfBeta * (b.smooth - white)) // RC Filter

		b.Out.Buffer[i] = finalScaling * b.smooth
	}
}
