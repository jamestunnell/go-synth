package units

type UnitCore interface {
	Configure(srate float64, p *Params) error
	NextSample(inputs []float64) []float64
}
