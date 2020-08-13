package processors_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit"
	"github.com/jamestunnell/go-synth/unit/processors"
	"github.com/stretchr/testify/assert"
)

func TestMulAddHappyPath(t *testing.T) {
	muladd := processors.MulAddPlugin.NewUnit()

	x := unit.NewBuffer(3)
	y := unit.NewBuffer(3)
	z := unit.NewBuffer(3)
	out := unit.NewBuffer(3)

	err := muladd.Initialize(100.0,
		map[string]*unit.Buffer{},
		[]*unit.Buffer{x, y, z}, []*unit.Buffer{out})

	if !assert.Nil(t, err) {
		return
	}

	x.Values[0] = 0.0
	x.Values[1] = 0.1
	x.Values[2] = 0.2

	y.Values[0] = 1.0
	y.Values[1] = -1.0
	y.Values[2] = 0.5

	z.Values[0] = 2.0
	z.Values[1] = -4.0
	z.Values[2] = 1.0

	muladd.Sample()

	assert.Equal(t, 2.0, out.Values[0])
	assert.Equal(t, -4.1, out.Values[1])
	assert.Equal(t, 1.1, out.Values[2])
}
