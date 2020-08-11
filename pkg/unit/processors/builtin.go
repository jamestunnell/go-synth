package processors

import "github.com/jamestunnell/go-synth/pkg/unit"

var (
	Builtin = []*unit.Plugin{
		AddKPlugin,
		AddXYPlugin,
		InvertPlugin,
		MulKPlugin,
		MulXYPlugin,
	}
)
