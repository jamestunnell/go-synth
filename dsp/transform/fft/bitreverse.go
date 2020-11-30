package fft

import (
	"fmt"
	"strconv"
	"strings"
)

func BitReverse(val uint64, nBits int) (uint64, error) {
	str := strconv.FormatUint(val, 2)
	nBitsInStr := len(str)

	if nBitsInStr > nBits {
		return 0, fmt.Errorf("%d requires more than %d bits", val, nBits)
	}

	str = ReverseString(str)
	// Add padding on the right, making the bit-reversed value left justified
	// since the original binary string should have been right-justified
	str += strings.Repeat("0", (nBits - nBitsInStr))
	val2, err := strconv.ParseUint(str, 2, nBits)
	if err != nil {
		return 0, fmt.Errorf("failed to parse bit reversed int string %s: %v", str, err)
	}

	return val2, nil
}
