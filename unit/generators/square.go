package generators

import (
	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/unit"
)

var (
	SquarePlugin = NewOscillatorPlugin(
		&unit.BasicInfo{
			Name:        "square",
			Description: "Naive square wave oscillator with 50% duty cycle from -1 to 1",
			Version:     "0.1.0-0",
			ID:          uuid.MustParse("17a5b4e3-e09c-4962-99d8-35906ac458b4"),
		},
		SquareWave)
)

func SquareWave(phase float64) float64 {
	var y float64
	if phase >= 0.0 {
		y = 1.0
	} else {
		y = -1.0
	}

	return y
}
