package generators

import (
	"math"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/unit"
)

const (
	// KTriangle is used to calculate a triangle wave
	KTriangle = 2.0 / math.Pi
)

var (
	TrianglePlugin = NewOscillatorPlugin(
		&unit.BasicInfo{
			Name:        "Triangle",
			Description: "Naive Triangle wave oscillator from -1 to 1",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("78455471-a278-4c20-9139-15f5389d6bdb"),
		},
		TriangleWave)
)

func TriangleWave(phase float64) float64 {
	return math.Abs(KTriangle*phase) - 1.0
}
