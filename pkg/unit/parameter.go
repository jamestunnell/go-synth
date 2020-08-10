package unit

// const (
// 	NameFieldName         = "name"
// 	DescriptionFieldName  = "description"
// 	RestrictionsFieldName = "description"
// )

type Parameter struct {
	Name        string       `json:"name`
	Description string       `json:"description"`
	Required    bool         `json:"required"`
	Default     float64      `json:"default"`
	Constraints []Constraint `json:"constraints"`
}

// func (p *Parameter) MarshalJSON() ([]byte, error) {
// 	restrictionObjs := make([]interface{}, len(p.Restrictions))

// 	for i, r := range p.Restrictions {
// 		restrictionObjs[i] =
// 	}

// 	obj := map[string]interface{}{
// 		NameFieldName: p.Name,
// 		DescriptionFieldName: p.Description,
// 		RestrictionsFieldName: restrictionObjs,
// 	}

// 	return json.Marshal(obj)
// }

// func (p *Parameter) UnmarshalJSON([]byte) error {

// }
