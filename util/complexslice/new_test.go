package complexslice_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/util/complexslice"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	f := func(i int) complex128 { return complex(float64(i), 0.0) }
	vals := complexslice.New(3, f)
	expected := []complex128{
		complex(0.0, 0.0),
		complex(1.0, 0.0),
		complex(2.0, 0.0),
	}

	assert.Equal(t, expected, vals)
}
