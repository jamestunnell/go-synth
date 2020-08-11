package unit

type NewUnitFunc func() Unit
type GetInterfaceFunc func(srate float64) *Interface

// Plugin can be used to create instances of a certain unit kind.
type Plugin struct {
	// BasicInfo contains the minimum information about a plugin
	BasicInfo *BasicInfo
	// NewUnit function creates a new unit
	NewUnit NewUnitFunc
	// GetInterface creates a description of the unit interface
	GetInterface GetInterfaceFunc
	// ExtraInfo contains any additional plugin information (author, website, etc.)
	ExtraInfo map[string]string
}
