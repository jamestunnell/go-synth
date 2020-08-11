package generators

import "github.com/jamestunnell/go-synth/unit"

var (
	Builtin = []*unit.Plugin{
		SquarePlugin,
		SinePlugin,
		SawtoothPlugin,
		TrianglePlugin,
	}
)
