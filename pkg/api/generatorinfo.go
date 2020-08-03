package api

import (
	"github.com/google/uuid"
)

type GeneratorMetadata struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
}

type GeneratorInfo struct {
	GeneratorMetadata
}
