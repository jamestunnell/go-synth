package generators

import (
	"math"

	"github.com/google/uuid"

	"github.com/jamestunnell/go-synth/pkg/unit"
)

const (
	// KSineB is used to calculate an approximation of a sine wave
	KSineB = 4.0 / math.Pi
	// KSineC is used to calculate an approximation of a sine wave
	KSineC = -4.0 / (math.Pi * math.Pi)
	// KSineP is used to calculate an approximation of a sine wave
	KSineP = 0.225
)

var (
	SinePlugin = NewOscillatorPlugin(
		&unit.BasicInfo{
			Name:        "sine",
			Description: "Sine wave (approximation) oscillator from -1 to 1",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("4626cfb7-e64a-4279-aade-cee0bdfe3ae0"),
		},
		SineWave)
)

func SineWave(phase float64) float64 {
	y := KSineB*phase + KSineC*phase*math.Abs(phase)
	// for extra precision
	return KSineP*(y*math.Abs(y)-y) + y // Q * y + P * y * y.abs
}
