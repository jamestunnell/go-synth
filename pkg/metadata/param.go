package metadata

type Restriction int

const (
	StrictlyPositive Restriction = iota
	NyquistFrequencyLimited
)

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Param struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Required     bool          `json:"required"`
	Default      float64       `json:"default,omitempty"`
	Ranges       []Range       `json:"ranges,omitempty"`
	Restrictions []Restriction `json:"restrictions,omitempty"`
}
