package osc

type Square struct {
	*Osc
}

// NewSquare makes a sine wave oscillator.
func NewSquare() *Square {
	return &Square{
		Osc: New(squareWave),
	}
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
