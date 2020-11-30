package complexslice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/util/complexslice"
)

func TestMapEmptySlice(t *testing.T) {
	count := 0
	x := complexslice.Map([]complex128{}, func(v complex128) complex128 {
		count++
		return 0.0
	})

	assert.Len(t, x, 0)
	assert.Equal(t, count, 0)
}

func TestMap(t *testing.T) {
	inputs := []complex128{
		complex(1.1, 0.0),
		complex(2.2, 0.0),
		complex(3.3, 0.0),
	}
	f := func(v complex128) complex128 { return v + 1.0 }
	expected := []complex128{
		complex(2.1, 0.0),
		complex(3.2, 0.0),
		complex(4.3, 0.0),
	}
	actual := complexslice.Map(inputs, f)

	assert.Equal(t, actual, expected)
}
