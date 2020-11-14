package node

import "fmt"

// Interface defines the node interface which is made up of inputs and controls.
// Inputs are required, so only their names are neeeded. Controls are optional
// so their default values are needed in case any controls are omitted.
type Interface struct {
	// Inputs are the names of node inputs
	InputNames []string `json:"inputNames"`
	// ControlDefaults maps the names of node controls to their defaults, which
	// will be used in case a control is omitted.
	ControlDefaults map[string]float64 `json:"controlDefaults"`
}

// NewInterface returns an empty interface
func NewInterface() *Interface {
	return &Interface{
		InputNames:      []string{},
		ControlDefaults: map[string]float64{},
	}
}

// EnsureInputs ensures that each interface input exists in the given map.
// Panics if an input is missing.
func (ifc *Interface) EnsureInputs(inputs Map) {
	for _, name := range ifc.InputNames {
		if _, found := inputs[name]; !found {
			panic(fmt.Sprintf("missing required input %s", name))
		}
	}
}

// EnsureControls ensures that each interface control exists in the given map.
// Adds a const node using the default value for any missing control.
func (ifc *Interface) EnsureControls(controls Map) {
	for name, defaultVal := range ifc.ControlDefaults {
		if _, found := controls[name]; !found {
			controls[name] = NewConst(defaultVal)
		}
	}
}
