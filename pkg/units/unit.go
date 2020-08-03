package units

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
