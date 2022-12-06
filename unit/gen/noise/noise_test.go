package noise_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/noise"
)

func TestBrown(t *testing.T) {
	testNoiseGen(t, noise.NewBrown())
}

func TestPink(t *testing.T) {
	testNoiseGen(t, noise.NewPink())
}

func TestWhite(t *testing.T) {
	testNoiseGen(t, noise.NewWhite())
}

func testNoiseGen(t *testing.T, b synth.Block) {
	ifc := synth.GetInterface(b)

	assert.Len(t, ifc.Params, 1)
	require.Contains(t, ifc.Params, "Seed")

	seed := ifc.Params["Seed"]

	assert.Len(t, ifc.Outputs, 1)
	require.Contains(t, ifc.Outputs, "Out")

	out := ifc.Outputs["Out"]

	require.NoError(t, seed.SetValue(time.Now().UnixNano()))

	require.NoError(t, b.Initialize(100.0, 10000))

	b.Configure()
	b.Run()

	vals, ok := out.Buffer().([]float64)

	require.True(t, ok)

	for i := 0; i < len(vals); i++ {
		assert.GreaterOrEqual(t, vals[i], -1.0)
		assert.Less(t, vals[i], 1.0)
	}

	// hist := histogram.Hist(20, vals)
	// maxWidth := 5
	// _ = histogram.Fprint(os.Stdout, hist, histogram.Linear(maxWidth))
}
