package fft_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/dsp/transform/fft"
)

func TestBitReverseHappyPath(t *testing.T) {
	testVals := []uint64{
		0, 1, 5, 25, 150003, 12345678910}

	for _, val := range testVals {
		testBitReverseHappyPath(t, val, 64)
	}

	testBitReverseHappyPath(t, 15, 4)
	testBitReverseHappyPath(t, 16, 5)
	testBitReverseHappyPath(t, 31, 5)
	testBitReverseHappyPath(t, 32, 6)
	testBitReverseHappyPath(t, 63, 6)
	testBitReverseHappyPath(t, 64, 7)
	testBitReverseHappyPath(t, 127, 7)
}

func TestBitReverseNotEnoughBits(t *testing.T) {
	testBitReverseNotEnoughBits(t, 15, 3)
	testBitReverseNotEnoughBits(t, 16, 4)
	testBitReverseNotEnoughBits(t, 31, 4)
	testBitReverseNotEnoughBits(t, 32, 5)
	testBitReverseNotEnoughBits(t, 63, 5)
	testBitReverseNotEnoughBits(t, 64, 6)
	testBitReverseNotEnoughBits(t, 127, 6)
}

func testBitReverseHappyPath(t *testing.T, val uint64, nBits int) {
	name := fmt.Sprintf("%d_happy_path_%d_bits", val, nBits)

	t.Run(name, func(t *testing.T) {
		rev, err := fft.BitReverse(val, nBits)

		if !assert.NoError(t, err) {
			return
		}

		val2, err := fft.BitReverse(rev, nBits)

		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, val, val2)
	})
}

func testBitReverseNotEnoughBits(t *testing.T, val uint64, nBits int) {
	name := fmt.Sprintf("%d_not_enough_bits_%d_bits", val, nBits)

	t.Run(name, func(t *testing.T) {
		_, err := fft.BitReverse(val, nBits)

		assert.Error(t, err)
	})
}
