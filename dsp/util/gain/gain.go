package gain

import (
	"fmt"
	"math"
)

// DecibelToLinear converts gain in decibels to linear.
func DecibelToLinear(dB float64) float64 {
	return math.Pow(10.0, dB/20.0)
}

// LinearToDecibel converts linear gain to decibels.
func LinearToDecibel(lin float64) (float64, error) {
	if lin <= 0.0 {
		return 0.0, fmt.Errorf("linear gain %f is not positive", lin)
	}

	return 20.0 * math.Log10(lin), nil
}
