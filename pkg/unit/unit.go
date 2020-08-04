package unit

import (
	"github.com/jamestunnell/go-synth/pkg/metadata"
)

type Unit struct {
	configured bool
	core       UnitCore
}

func NewUnit(core UnitCore) *Unit {
	return &Unit{
		configured: false,
		core:       core,
	}
}

func (u *Unit) Metadata() *metadata.Metadata {
	return u.core.Metadata()
}

func (u *Unit) IsConfigured() bool {
	return u.configured
}

func (u *Unit) Configure(srate float64, p *Params) (bool, error) {
	if u.configured {
		return false, nil
	}

	err := u.core.Configure(srate, p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *Unit) FillInDefaultParams(p *Params) {
	m := u.core.Metadata()
	for _, param := range m.Parameters {
		if !p.HasParamValue(param.Name) && !param.Required {
			p.ParamValues[param.Name] = param.Default
		}
	}
}

func (u *Unit) NextSample(inputs []float64) []float64 {
	return u.core.NextSample(inputs)
}