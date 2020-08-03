package metadata

import "github.com/google/uuid"

type Metadata struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Author      string            `json:"author,omitempty"`
	Parameters  []Param           `json:"paramters,omitempty"`
	Inputs      []NameDescription `json:"inputs,omitempty"`
	Outputs     []NameDescription `json:"outputs,omitempty"`
}
