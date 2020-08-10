package constraints

const typeStrNotInSet = "NotInSet"

type NotInSet struct {
	MultiValue
}

func NewNotInSet(values []float64) *NotInSet {
	return &NotInSet{MultiValue{Values: values, Type: typeStrNotInSet}}
}

func (notInSet *NotInSet) Allows(val float64) bool {
	for _, v := range notInSet.Values {
		if v == val {
			return false
		}
	}

	return true
}
