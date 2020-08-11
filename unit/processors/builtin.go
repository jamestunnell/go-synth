package processors

import "github.com/jamestunnell/go-synth/unit"

var (
	Builtin = []*unit.Plugin{
		AddKPlugin,
		AddXYPlugin,
		InvertPlugin,
		MulKPlugin,
		MulXYPlugin,
	}
)
