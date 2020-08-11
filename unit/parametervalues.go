package unit

type ParameterValues struct {
	NameIndexMap map[string]int
	Values       []float64
}

func NewParameterValues(nameValueMap map[string]float64) *ParameterValues {
	nameIndexMap := make(map[string]int)
	values := make([]float64, len(nameValueMap))

	i := 0
	for name, value := range nameValueMap {
		values[i] = value
		nameIndexMap[name] = i

		i++
	}

	return &ParameterValues{NameIndexMap: nameIndexMap, Values: values}
}

func (p *ParameterValues) GetNameIndices(names []string) ([]int, bool) {
	indices := make([]int, len(names))

	for i, name := range names {
		index, found := p.NameIndexMap[name]
		if !found {
			return []int{}, false
		}

		indices[i] = index
	}

	return indices, true
}

// func (p *ParameterValues) HasValue(paramName string) bool {
// 	_, found := p.Values[paramName]
// 	return found
// }

// func (p *ParameterValues) GetValue(paramName string) (float64, error) {
// 	val, found := p.Values[paramName]
// 	if found {
// 		return 0.0, fmt.Errorf("param not found: %s", paramName)
// 	}

// 	return val, nil
// }
