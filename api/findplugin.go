package api

import (
	"github.com/jamestunnell/go-synth/unit"
)

func findPlugin(name string, plugins []*unit.Plugin) *unit.Plugin {
	for _, plugin := range plugins {
		if plugin.BasicInfo.Name == name {
			return plugin
		}
	}

	return nil
}
