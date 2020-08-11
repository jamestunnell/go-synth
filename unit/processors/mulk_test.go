package processors_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/unit"
	"github.com/jamestunnell/go-synth/unit/processors"
	"github.com/stretchr/testify/assert"
)

func TestMulKHappyPath(t *testing.T) {
	mulK := processors.MulKPlugin.NewUnit()

	x := unit.NewBuffer(3)
	k := unit.NewBuffer(1)
	z := unit.NewBuffer(3)

	err := mulK.Initialize(100.0,
		map[string]*unit.Buffer{processors.ParamNameK: k},
		[]*unit.Buffer{x}, []*unit.Buffer{z})

	if !assert.Nil(t, err) {
		return
	}

	k.Values[0] = 2.0

	mulK.Configure()

	x.Values[0] = 0.0
	x.Values[1] = 0.1
	x.Values[2] = 0.2

	mulK.Sample()

	assert.Equal(t, 0.0, z.Values[0])
	assert.Equal(t, 0.2, z.Values[1])
	assert.Equal(t, 0.4, z.Values[2])
}
