package units

import (
	"math"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/pkg/metadata"
)

type SquareWave struct {
	phaseIncr float64
	phase     float64
}

const (
	DefaultPhaseOffset   = 0.0
	ParamNameFreq        = "freq"
	ParamNamePhaseOffset = "phaseOffset"
	TwoPi                = 2.0 * math.Pi
)

var (
	SquareWaveMetadata = metadata.Metadata{
		ID:          uuid.MustParse("9c575a62-8d7a-46c5-ab72-92d342f3e238"),
		Name:        "squarewave",
		Description: "square wave oscillator with 50% duty cycle from -1 to 1",
		Author:      "James Tunnell",
		Parameters: []metadata.Param{
			{
				Name:        ParamNameFreq,
				Description: "wave frequency",
				Required:    true,
				Restrictions: []metadata.Restriction{
					metadata.NyquistFrequencyLimited,
					metadata.StrictlyPositive,
				},
			},
			{
				Name:        ParamNamePhaseOffset,
				Description: "wave phase offset",
				Required:    false,
				Default:     DefaultPhaseOffset,
				Ranges:      []metadata.Range{{Min: -TwoPi, Max: TwoPi}},
			},
		},
		Outputs: []metadata.NameDescription{{Name: "out1"}},
	}
)

func (sq *SquareWave) Configure(srate float64, p *Params) error {
	freq, err := p.GetParamValue(ParamNameFreq)
	if err != nil {
		return err
	}

	phaseOffset, err := p.GetParamValue(ParamNamePhaseOffset)
	if err != nil {
		return err
	}

	sq.phaseIncr = (freq * TwoPi) / srate
	sq.phase = phaseOffset

	return nil
}

func (sq *SquareWave) NextSample(inputs []float64) []float64 {
	var x float64
	if sq.phase >= 0.0 {
		x = 1.0
	} else {
		x = -1.0
	}

	sq.phase += sq.phaseIncr
	if sq.phase >= math.Pi {
		sq.phase -= TwoPi
	}

	return []float64{x}
}
