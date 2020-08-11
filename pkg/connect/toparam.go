package connect

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/jamestunnell/go-synth/pkg/unit"
)

type ToParam struct {
	ParamName     string
	TargetID uuid.UUID
}

func (to *ToParam) ConnectedBuffer(wrappers map[uuid.UUID]*UnitWrapper) (*unit.Buffer, error) {
	targetWrapper, found := wrappers[to.TargetID]
	if !found {
		return nil, fmt.Errorf("unit wrapper with ID %v not found", to.TargetID)
	}

	paramBuf, found := targetWrapper.ParamBuffers[to.ParamName]
	if !found {
		return nil, fmt.Errorf("param buffer with name %s not found", to.ParamName)
	}

	return paramBuf, nil
}
