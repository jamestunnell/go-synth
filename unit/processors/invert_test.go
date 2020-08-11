package processors_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit"
	"github.com/jamestunnell/go-synth/unit/processors"
	"github.com/stretchr/testify/assert"
)

func TestInvertHappyPath(t *testing.T) {
	inv := processors.InvertPlugin.NewUnit()

	x := unit.NewBuffer(3)
	z := unit.NewBuffer(3)

	err := inv.Initialize(100.0, map[string]*unit.Buffer{},
		[]*unit.Buffer{x}, []*unit.Buffer{z})

	if !assert.Nil(t, err) {
		return
	}

	x.Values[0] = 1.0
	x.Values[1] = 0.5
	x.Values[2] = -0.5

	inv.Sample()

	assert.Equal(t, 1.0, z.Values[0])
	assert.Equal(t, 2.0, z.Values[1])
	assert.Equal(t, -2.0, z.Values[2])
}
