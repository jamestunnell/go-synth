package generators

import "github.com/jamestunnell/go-synth/pkg/unit"

var (
	Builtin = []*unit.Plugin{
		SquarePlugin,
		SinePlugin,
		SawtoothPlugin,
		TrianglePlugin,
	}
)
