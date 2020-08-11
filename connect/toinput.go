package connect

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/jamestunnell/go-synth/unit"
)

type ToInput struct {
	InputIndex int
	TargetID   uuid.UUID
}

func (to *ToInput) ConnectedBuffer(wrappers map[uuid.UUID]*UnitWrapper) (*unit.Buffer, error) {
	targetWrapper, found := wrappers[to.TargetID]
	if !found {
		return nil, fmt.Errorf("unit wrapper with ID %v not found", to.TargetID)
	}

	n := len(targetWrapper.InBuffers)

	if to.InputIndex >= n {
		return nil, fmt.Errorf("input index %d is not in valid range [0,%d)", to.InputIndex, n)
	}

	return targetWrapper.InBuffers[to.InputIndex], nil
}
