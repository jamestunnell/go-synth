package node

import (
	"fmt"

	"github.com/jamestunnell/go-synth/util/param"
)

// Interface defines the node interface which is made up of inputs and controls.
// Inputs are required, so only their names are neeeded. Controls are optional
// so their default values are needed in case any controls are omitted.
type Interface struct {
	// Inputs are the names of node inputs
	InputNames []string `json:"inputNames"`
	// ControlDefaults maps the names of node controls to their defaults, which
	// will be used in case a control is omitted.
	ControlDefaults map[string]float64    `json:"controlDefaults"`
	ParamTypes      map[string]param.Type `json:"paramTypes"`
}

// NewInterface returns an empty interface
func NewInterface() *Interface {
	return &Interface{
		InputNames:      []string{},
		ControlDefaults: map[string]float64{},
		ParamTypes:      map[string]param.Type{},
	}
}

// CheckInputs ensures that each interface input exists in the given map.
// Panics if an input is missing.
func (ifc *Interface) CheckInputs(inputs Map) error {
	for _, name := range ifc.InputNames {
		if _, found := inputs[name]; !found {
			return fmt.Errorf("missing required input %s", name)
		}
	}

	return nil
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

// CheckParams ensures that each interface param exists in the given map.
// Panics if a param is missing or if it is not the expected type.
func (ifc *Interface) CheckParams(params param.Map) error {
	for name, paramType := range ifc.ParamTypes {
		if p, found := params[name]; !found {
			return fmt.Errorf("missing required param %s", name)
		} else if p.Type() != paramType {
			return fmt.Errorf("param %s is not type %s", name, paramType)
		}
	}

	return nil
}
