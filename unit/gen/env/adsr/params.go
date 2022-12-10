package adsr

import (
	"fmt"
)

// Params is used to carry the parameter values needed for ADSR.
type Params struct {
	SustainLevel, AttackTime, DecayTime, ReleaseTime float64
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
