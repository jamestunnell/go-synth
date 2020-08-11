package processors_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/pkg/unit"
	"github.com/jamestunnell/go-synth/pkg/unit/processors"
	"github.com/stretchr/testify/assert"
)

func TestAddKHappyPath(t *testing.T) {
	addK := processors.AddKPlugin.NewUnit()

	x := unit.NewBuffer(3)
	k := unit.NewBuffer(1)
	z := unit.NewBuffer(3)

	err := addK.Initialize(100.0,
		map[string]*unit.Buffer{processors.ParamNameK: k},
		[]*unit.Buffer{x}, []*unit.Buffer{z})

	if !assert.Nil(t, err) {
		return
	}

	k.Values[0] = 1.0

	addK.Configure()

	x.Values[0] = 0.0
	x.Values[1] = 0.1
	x.Values[2] = 0.2

	addK.Sample()

	assert.Equal(t, 1.0, z.Values[0])
	assert.Equal(t, 1.1, z.Values[1])
	assert.Equal(t, 1.2, z.Values[2])
}
