package connect

import (
	"github.com/google/uuid"

	"github.com/jamestunnell/go-synth/unit"
)

type Connection interface {
	ConnectedBuffer(map[uuid.UUID]*UnitWrapper) (*unit.Buffer, error)
}