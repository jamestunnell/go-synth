package adsr

import (
	"fmt"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/mod"
	"github.com/jamestunnell/go-synth/util/param"
)

// Params is used to carry the parameter values needed for ADSR.
type Params struct {
	PeakLevel, SustainLevel, AttackMs, DecayMs, ReleaseMs float64
}

// NewParamsFromMap makes a new Params instance using the given parameter map.
func NewParamsFromMap(m param.Map) *Params {
	p := &Params{}

	p.PeakLevel = m[ParamNamePeakLevel].Value().(float64)
	p.SustainLevel = m[ParamNamePeakLevel].Value().(float64)
	p.AttackMs = m[ParamNameAttackMs].Value().(float64)
	p.DecayMs = m[ParamNameDecayMs].Value().(float64)
	p.ReleaseMs = m[ParamNameReleaseMs].Value().(float64)

	return p
}

// MakeMods constructs mod functions for adding the parameters to the node.
func (p *Params) MakeMods() []node.Mod {
	return []node.Mod{
		mod.Param(ParamNamePeakLevel, param.NewFloat(p.PeakLevel)),
		mod.Param(ParamNameSustainLevel, param.NewFloat(p.SustainLevel)),
		mod.Param(ParamNameAttackMs, param.NewFloat(p.AttackMs)),
		mod.Param(ParamNameDecayMs, param.NewFloat(p.DecayMs)),
		mod.Param(ParamNameReleaseMs, param.NewFloat(p.ReleaseMs)),
	}
}

// Validate checks the parameters values.
// Peak level must be positive, and sustain level must  be non-negative.
// Peak level must not be less than sustain level.
// Attack, decay, and release times must be non-negative.
// Returns non-nil error if any parameter is invalid.
func (p *Params) Validate() error {
	if p.PeakLevel <= 0.0 {
		return fmt.Errorf("peak level %f is not positive", p.PeakLevel)
	}

	if p.SustainLevel < 0.0 {
		return fmt.Errorf("sustain level %f is negative", p.SustainLevel)
	}

	if p.PeakLevel < p.SustainLevel {
		return fmt.Errorf("peak level %f is less than sustain level %f", p.PeakLevel, p.SustainLevel)
	}

	if p.AttackMs <= 0.0 {
		return fmt.Errorf("attack time msec %f is not positive", p.AttackMs)
	}

	if p.DecayMs <= 0.0 {
		return fmt.Errorf("decay time msec %f is not positive", p.DecayMs)
	}

	if p.ReleaseMs <= 0.0 {
		return fmt.Errorf("release time msec %f is not positive", p.ReleaseMs)
	}

	return nil
}
