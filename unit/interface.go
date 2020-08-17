package unit

import "errors"

type Interface struct {
	Parameters map[string]*ParamInfo `json:"parameters"`
	NumInputs  int                   `json:"numInputs"`
	NumOutputs int                   `json:"numOutputs"`
}

func (ifc *Interface) Verify() error {
	if ifc.NumInputs == 0 && ifc.NumOutputs == 0 {
		return errors.New("interface has no inputs and no outputs")
	}

	return nil
}
