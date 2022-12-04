package brown

import (
	"fmt"

	"github.com/jamestunnell/go-synth/unit/gen/noise/white"
)

// Brown produces brown noise by running white noise through a lowpass filter.
// Adapted from https://github.com/alessandrocuda/noise_generator
// Output is from -1 to 1.
type Brown struct {
	*white.White
	smooth float64
	outBuf []float64
}

// scaling is applied to make the final output range close to [-1,1)
const finalScaling = 3.0

const lpfBeta = 0.025

// New makes a new brown noise block.
func New() *Brown {
	return &Brown{
		White:  white.New(),
		smooth: 0,
		outBuf: []float64{},
	}
}

// Initialize initializes the node.
func (b *Brown) Initialize(srate float64, outDepth int) error {
	if err := b.White.Initialize(srate, outDepth); err != nil {
		return fmt.Errorf("failed to init white noise: %w", err)
	}

	b.smooth = 0
	b.outBuf = b.White.Out.Buffer().([]float64)

	return nil
}

// Run runs the brown noise generation process.
func (b *Brown) Run() {
	// generate the white noise
	b.White.Run()

	for i := 0; i < len(b.outBuf); i++ {
		white := b.outBuf[i]

		b.smooth = b.smooth - (lpfBeta * (b.smooth - white)) // RC Filter

		b.outBuf[i] = finalScaling * b.smooth
	}
}
