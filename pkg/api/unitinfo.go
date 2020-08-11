package api

import (
	"github.com/jamestunnell/go-synth/pkg/unit"
)

type UnitInfo struct {
	BasicInfo *unit.BasicInfo   `json:"basicInfo"`
	Interface *unit.Interface   `json:"interface"`
	ExtraInfo map[string]string `json:"extraInfo,omitempty"`
}
