package osc

func Square(params *Params) *Osc {
	return new(params, squareWave)
}

func squareWave(phase float64) float64 {
	var y float64
	if phase >= 0.0 {
		y = 1.0
	} else {
		y = -1.0
	}

	return y
}
