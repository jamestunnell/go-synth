package api

import (
	"strings"

	"github.com/jamestunnell/go-synth/node"
)

func findCore(name string, reg *node.CoreRegistry) node.Core {
	for _, path := range reg.Paths() {
		if strings.Contains(path, name) {
			if core, ok := reg.GetCore(path); ok {
				return core
			}
		}
	}

	return nil
}
