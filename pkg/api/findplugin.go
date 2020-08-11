package api

import (
	"github.com/jamestunnell/go-synth/pkg/unit"
	"github.com/jamestunnell/go-synth/pkg/unit/generators"
)

func findGeneratorPlugin(name string) *unit.Plugin {
	for _, plugin := range generators.Builtin {
		if plugin.BasicInfo.Name == name {
			return plugin
		}
	}

	return nil
}
