package constraints_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jamestunnell/go-synth/pkg/unit/constraints"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalSingleValueConstraint(t *testing.T) {
	testUnmarshalConstraintJSON(t, constraints.NewLess(5))
	testUnmarshalConstraintJSON(t, constraints.NewGreater(5))
	testUnmarshalConstraintJSON(t, constraints.NewLessEqual(5))
	testUnmarshalConstraintJSON(t, constraints.NewGreaterEqual(5))
	testUnmarshalConstraintJSON(t, constraints.NewEqual(5))
}

func TestUnmarshalMultiValueConstraint(t *testing.T) {
	values := []float64{0.0, -1.5, 3.42}

	testUnmarshalConstraintJSON(t, constraints.NewInSet(values))
	testUnmarshalConstraintJSON(t, constraints.NewNotInSet(values))
}

func testUnmarshalConstraintJSON(t *testing.T, constraint1 interface{}) {
	data, err := json.Marshal(constraint1)

	if !assert.Nil(t, err) {
		return
	}

	constraint2, err := constraints.UnmarshalConstraintJSON(data)

	if !assert.Nil(t, err) {
		return
	}

	assert.Equal(t, reflect.TypeOf(constraint1), reflect.TypeOf(constraint2))
}
