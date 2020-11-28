package decay

import "math"

// TargetDecay is the target envelope value to reach after decay
// time has passed since trigger
const TargetDecay = 0.01

// Multiplier calculates the multiplier to use in an exponential
// decay process, where the starting value is 1.0 and the target
// end value after n multiplications is TargetDecay.
func Multiplier(n int) float64 {
	return math.Pow(10.0, math.Log10(TargetDecay)/float64(n))
}
