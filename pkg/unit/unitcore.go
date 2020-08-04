package unit

import (
	"github.com/jamestunnell/go-synth/pkg/metadata"
)

type UnitCore interface {
	New() UnitCore
	Metadata() *metadata.Metadata
	Configure(srate float64, p *Params) error
	NextSample(inputs []float64) []float64
}
