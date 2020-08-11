package constraints

const typeStrInSet = "InSet"

type InSet struct {
	MultiValue
}

func NewInSet(values []float64) *InSet {
	return &InSet{MultiValue{Values: values, Type: typeStrInSet}}
}

func (inSet *InSet) Allows(val float64) bool {
	for _, v := range inSet.Values {
		if v == val {
			return false
		}
	}

	return true
}
