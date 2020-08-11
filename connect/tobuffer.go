package connect

import (
	"github.com/google/uuid"

	"github.com/jamestunnell/go-synth/unit"
)

type ToBuffer struct {
	TargetBuffer *unit.Buffer
}

func (to *ToBuffer) ConnectedBuffer(map[uuid.UUID]*UnitWrapper) (*unit.Buffer, error) {
	return to.TargetBuffer, nil
}
