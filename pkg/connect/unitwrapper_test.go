package connect_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/pkg/connect"
	"github.com/jamestunnell/go-synth/pkg/unit"
	"github.com/jamestunnell/go-synth/pkg/unit/processors"
	"github.com/stretchr/testify/assert"
)

const srate = 100.0

func TestUnitWrapperNewUnconnected(t *testing.T) {
	u := processors.AddKPlugin.NewUnit()
	ifc := processors.AddKPlugin.GetInterface(srate)
	w := connect.NewUnitWrapper(u, ifc, 10)
	wrappers := map[uuid.UUID]*connect.UnitWrapper{uuid.New(): w}

	assert.Equal(t, ifc.NumInputs, len(w.InBuffers))
	assert.Equal(t, ifc.NumOutputs, len(w.OutConnections))
	assert.Equal(t, len(ifc.Parameters), len(w.ParamBuffers))

	// should fail because output is not connected
	err := w.InitializeUnit(wrappers, srate)

	assert.NotNil(t, err)
}

func TestUnitWrapperConnectedToMissingWrapper(t *testing.T) {
	w := makeUnitWrapper(processors.AddKPlugin, 10)
	wrappers := map[uuid.UUID]*connect.UnitWrapper{uuid.New(): w}

	w.OutConnections[0] = &connect.ToInput{TargetID: uuid.New(), InputIndex: 0}

	err := w.InitializeUnit(wrappers, srate)

	assert.NotNil(t, err)
}

func TestUnitWrapperConnectsToInput(t *testing.T) {
	w1 := makeUnitWrapper(processors.AddKPlugin, 10)
	w2 := makeUnitWrapper(processors.MulKPlugin, 10)
	id2 := uuid.New()
	wrappers := map[uuid.UUID]*connect.UnitWrapper{uuid.New(): w1, id2: w2}

	w1.OutConnections[0] = &connect.ToInput{TargetID: id2, InputIndex: 0}

	err := w1.InitializeUnit(wrappers, srate)

	assert.Nil(t, err)

	// Also test that the output can't be connected to a param
	// because the buffer depth is > 1
	w1.OutConnections[0] = &connect.ToParam{TargetID: id2, ParamName: processors.ParamNameK}

	err = w1.InitializeUnit(wrappers, srate)

	assert.NotNil(t, err)
}

func TestUnitWrapperConnectsToParam(t *testing.T) {
	w1 := makeUnitWrapper(processors.AddKPlugin, 1)
	w2 := makeUnitWrapper(processors.MulKPlugin, 10)
	id2 := uuid.New()
	wrappers := map[uuid.UUID]*connect.UnitWrapper{uuid.New(): w1, id2: w2}

	w1.OutConnections[0] = &connect.ToParam{TargetID: id2, ParamName: processors.ParamNameK}

	err := w1.InitializeUnit(wrappers, srate)

	assert.Nil(t, err)
}

func makeUnitWrapper(plugin *unit.Plugin, bufferDepth int) *connect.UnitWrapper {
	u := plugin.NewUnit()
	ifc := plugin.GetInterface(srate)
	return connect.NewUnitWrapper(u, ifc, bufferDepth)
}
