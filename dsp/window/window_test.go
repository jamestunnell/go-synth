package window_test

import (
	"testing"

	"github.com/jamestunnell/go-synth/dsp/window"
	"github.com/stretchr/testify/assert"
)

func TestBartlett(t *testing.T) {
	testWindowMaker(t, window.NewBartlett())
}

func TestBartlettHann(t *testing.T) {
	testWindowMaker(t, window.NewBartlettHann())
}

func TestBlackman(t *testing.T) {
	testWindowMaker(t, window.NewBlackman())
}

func TestBlackmanHarris(t *testing.T) {
	testWindowMaker(t, window.NewBlackmanHarris())
}

func TestBlackmanNuttall(t *testing.T) {
	testWindowMaker(t, window.NewBlackmanNuttall())
}

func TestCosine(t *testing.T) {
	testWindowMaker(t, window.NewCosine())
}

func TestFlatTop(t *testing.T) {
	testWindowMaker(t, window.NewFlatTop())
}

func TestGaussian(t *testing.T) {
	_, err := window.NewGaussian(0.0)
	assert.NoError(t, err)

	w, err := window.NewGaussian(0.5)
	assert.NoError(t, err)
	testWindowMaker(t, w)
}

func TestGaussianBadSigma(t *testing.T) {
	_, err := window.NewGaussian(0.51)
	assert.Error(t, err)

	_, err = window.NewGaussian(-0.01)
	assert.Error(t, err)
}

func TestHamming(t *testing.T) {
	testWindowMaker(t, window.NewHamming())
}

func TestHann(t *testing.T) {
	testWindowMaker(t, window.NewHann())
}

func TestLanczos(t *testing.T) {
	testWindowMaker(t, window.NewLanczos())
}

func TestNuttall(t *testing.T) {
	testWindowMaker(t, window.NewNuttall())
}

func TestRectangular(t *testing.T) {
	testWindowMaker(t, window.NewRectangular())
}

func TestTriangular(t *testing.T) {
	testWindowMaker(t, window.NewTriangular())
}

func TestTukey(t *testing.T) {
	_, err := window.NewTukey(0.0)
	assert.NoError(t, err)

	w, err := window.NewTukey(1.0)
	assert.NoError(t, err)

	testWindowMaker(t, w)
}

func TestTukeyBadAlpha(t *testing.T) {
	_, err := window.NewTukey(-0.01)
	assert.Error(t, err)

	_, err = window.NewTukey(1.01)
	assert.Error(t, err)
}

func testWindowMaker(t *testing.T, w window.WindowMaker) {
	assert.Len(t, w.Make(1), 1)
	assert.Len(t, w.Make(4), 4)
	assert.Len(t, w.Make(47), 47)
}
