package unit

import "errors"

type NewUnitFunc func() Unit

// Plugin can be used to create instances of a certain unit kind.
type Plugin struct {
	// BasicInfo contains the minimum information about a plugin
	BasicInfo *BasicInfo `json:"basicInfo"`
	// Interface describes the unit interface
	Interface *Interface `json:"interface"`
	// NewUnit function creates a new unit
	NewUnit NewUnitFunc `json:"-"`
	// ExtraInfo contains any additional plugin information (author, website, etc.)
	ExtraInfo map[string]string `json:"extraInfo,omitempty"`
}

func (p *Plugin) Verify() error {
	if p.BasicInfo != nil {
		err := p.BasicInfo.Verify()
		if err != nil {
			return err
		}
	} else {
		return errors.New("basic info not given")
	}

	if p.Interface != nil {
		err := p.Interface.Verify()
		if err != nil {
			return err
		}
	} else {
		return errors.New("interface not given")
	}

	if p.NewUnit == nil {
		return errors.New("new unit func not given")
	}

	return nil
}
