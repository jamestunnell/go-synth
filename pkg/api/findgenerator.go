package api

import (
	"github.com/jamestunnell/go-synth/pkg/unit"
	"github.com/jamestunnell/go-synth/pkg/unit/generators"
)

func findGenerator(name string) unit.UnitCore {
	for _, core := range generators.GeneratorBlanks {
		if core.Metadata().Name == name {
			return core.New()
		}
	}

	return nil
}