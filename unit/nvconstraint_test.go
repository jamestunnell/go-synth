package unit_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit"
	"github.com/stretchr/testify/assert"
)

func TestNVConstraint(t *testing.T) {
	all := []unit.NVConstraintType{
		unit.Integer,
		unit.Positive,
		unit.NonNegative,
		unit.NyquistLimited,
	}

	for _, t1 := range all {
		i := t1.Info()
		t2 := i.Type()

		assert.Equal(t, t1, t2)
	}
}
