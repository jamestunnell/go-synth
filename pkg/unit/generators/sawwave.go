package generators

import (
	"math"

	"github.com/jamestunnell/go-synth/pkg/metadata"
	"github.com/jamestunnell/go-synth/pkg/unit"
)

type SawWave struct {
	phaseIncr float64
	phase     float64
}

var (
	SawWaveMetadata = &metadata.Metadata{
		Name:        "saw",
		Description: "Sawtooth wave oscillator from -1 to 1",
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
		NumOutputs: 1,
	}
)

func (sq *SawWave) New() unit.UnitCore {
	return &SawWave{}
}

func (sq *SawWave) Metadata() *metadata.Metadata {
	return SawWaveMetadata
}

func (sq *SawWave) Configure(srate float64, p *unit.Params) error {
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

func (sq *SawWave) NextSample(inputs []float64) []float64 {
	x := sq.phase * OverPi

	sq.phase += sq.phaseIncr
	if sq.phase >= math.Pi {
		sq.phase -= TwoPi
	}

	return []float64{x}
}
