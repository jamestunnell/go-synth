package unit

type ParamInfo struct {
	Description string       `json:"description"`
	Required    bool         `json:"required"`
	Default     float64      `json:"default"`
	Constraints []Constraint `json:"constraints"`
}
