package connect

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/jamestunnell/go-synth/pkg/unit"
)

type UnitWrapper struct {
	Unit           unit.Unit
	BufferDepth    int
	ParamBuffers   map[string]*unit.Buffer
	InBuffers      []*unit.Buffer
	OutConnections []Connection
}

func NewUnitWrapper(u unit.Unit, ifc *unit.Interface, bufDepth int) *UnitWrapper {
	inBuffers := make([]*unit.Buffer, ifc.NumInputs)
	outConnections := make([]Connection, ifc.NumOutputs)
	paramBuffers := make(map[string]*unit.Buffer)

	for i := 0; i < ifc.NumInputs; i++ {
		inBuffers[i] = unit.NewBuffer(bufDepth)
	}

	for name := range ifc.Parameters {
		paramBuffers[name] = unit.NewBuffer(1)
	}

	return &UnitWrapper{
		Unit:           u,
		BufferDepth:    bufDepth,
		ParamBuffers:   paramBuffers,
		InBuffers:      inBuffers,
		OutConnections: outConnections}
}

func (w *UnitWrapper) InitializeUnit(wrappers map[uuid.UUID]*UnitWrapper, srate float64) error {
	outBuffers := make([]*unit.Buffer, len(w.OutConnections))

	unconnected := []int{}
	for i, conn := range w.OutConnections {
		if conn == nil {
			unconnected = append(unconnected, i)
		}
	}

	if len(unconnected) > 0 {
		return fmt.Errorf("outputs %v are not connected", unconnected)
	}

	for i, conn := range w.OutConnections {
		outBuf, err := conn.ConnectedBuffer(wrappers)
		if outBuf == nil {
			return fmt.Errorf("failed to connect output %d: %v", i, err)
		}

		if outBuf.Length != w.BufferDepth {
			return fmt.Errorf("output buffer depth %d is not %d", outBuf.Length, w.BufferDepth)
		}

		outBuffers[i] = outBuf
	}

	return w.Unit.Initialize(srate, w.ParamBuffers, w.InBuffers, outBuffers)
}
