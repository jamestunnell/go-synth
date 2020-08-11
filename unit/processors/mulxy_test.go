package processors_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit"
	"github.com/jamestunnell/go-synth/unit/processors"
	"github.com/stretchr/testify/assert"
)

func TestMulXYHappyPath(t *testing.T) {
	mulXY := processors.MulXYPlugin.NewUnit()

	x := unit.NewBuffer(3)
	y := unit.NewBuffer(3)
	z := unit.NewBuffer(3)

	err := mulXY.Initialize(100.0, map[string]*unit.Buffer{},
		[]*unit.Buffer{x, y}, []*unit.Buffer{z})

	if !assert.Nil(t, err) {
		return
	}

	x.Values[0] = 0.0
	x.Values[1] = 0.1
	x.Values[2] = 0.2

	y.Values[0] = 1.0
	y.Values[1] = -1.0
	y.Values[2] = 1.0

	mulXY.Sample()

	assert.Equal(t, 0.0, z.Values[0])
	assert.Equal(t, -0.1, z.Values[1])
	assert.Equal(t, 0.2, z.Values[2])
}
