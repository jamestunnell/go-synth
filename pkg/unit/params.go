package unit

import (
	"fmt"
)

type Params struct {
	ParamValues map[string]float64
}

func NewParams(paramValues map[string]float64) *Params {
	return &Params{ParamValues: paramValues}
}

func (p *Params) HasParamValue(paramName string) bool {
	_, found := p.ParamValues[paramName]
	return found
}

func (p *Params) GetParamValue(paramName string) (float64, error) {
	val, found := p.ParamValues[paramName]
	if !found {
		return 0.0, fmt.Errorf("param not found: %s", paramName)
	}

	return val, nil
}
