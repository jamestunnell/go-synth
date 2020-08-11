package generators

import (
	"math"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/unit"
)

const OneOverPi = 1.0 / math.Pi

var (
	SawtoothPlugin = NewOscillatorPlugin(
		&unit.BasicInfo{
			Name:        "sawtooth",
			Description: "Naive sawtooth wave oscillator from -1 to 1",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("c477ef01-3ebf-4162-bcb2-b9df98cb6dfb"),
		},
		SawtoothWave)
)

func SawtoothWave(phase float64) float64 {
	return phase * OneOverPi
}
