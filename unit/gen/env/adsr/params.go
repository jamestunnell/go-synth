package adsr

import (
	"fmt"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/mod"
	"github.com/jamestunnell/go-synth/util/param"
)

// Params is used to carry the parameter values needed for ADSR.
type Params struct {
	SustainLevel, AttackTime, DecayTime, ReleaseTime float64
}

// NewParamsFromMap makes a new Params instance using the given parameter map.
func NewParamsFromMap(m param.Map) *Params {
	p := &Params{}

	p.SustainLevel = m[ParamNameSustainLevel].Value().(float64)
	p.AttackTime = m[ParamNameAttackTime].Value().(float64)
	p.DecayTime = m[ParamNameDecayTime].Value().(float64)
	p.ReleaseTime = m[ParamNameReleaseTime].Value().(float64)

	return p
}

// MakeMods constructs mod functions for adding the parameters to the node.
func (p *Params) MakeMods() []node.Mod {
	return []node.Mod{
		mod.Param(ParamNameSustainLevel, param.NewFloat(p.SustainLevel)),
		mod.Param(ParamNameAttackTime, param.NewFloat(p.AttackTime)),
		mod.Param(ParamNameDecayTime, param.NewFloat(p.DecayTime)),
		mod.Param(ParamNameReleaseTime, param.NewFloat(p.ReleaseTime)),
	}
}

// Validate checks the parameters values.
// Peak level must be positive, and sustain level must  be non-negative.
// Peak level must not be less than sustain level.
// Attack, decay, and release times must be non-negative.
// Returns non-nil error if any parameter is invalid.
func (p *Params) Validate() error {
	if p.SustainLevel < 0.0 {
		return fmt.Errorf("sustain level %f is negative", p.SustainLevel)
	}

	if p.SustainLevel > 1.0 {
		return fmt.Errorf("sustain level %f is greater than 1", p.SustainLevel)
	}

	if p.AttackTime <= 0.0 {
		return fmt.Errorf("attack time %f is not positive", p.AttackTime)
	}

	if p.DecayTime <= 0.0 {
		return fmt.Errorf("decay time %f is not positive", p.DecayTime)
	}

	if p.ReleaseTime <= 0.0 {
		return fmt.Errorf("release time %f is not positive", p.ReleaseTime)
	}

	return nil
}
