package connect_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/connect"
	"github.com/jamestunnell/go-synth/unit"
	"github.com/stretchr/testify/assert"
)

func TestToInput(t *testing.T) {
	const inputIndex = 1

	id := uuid.New()
	wrappers := map[uuid.UUID]*connect.UnitWrapper{}
	toInput := connect.ToInput{TargetID: id, InputIndex: inputIndex}
	ifcWithoutInput := &unit.Interface{
		Parameters: map[string]*unit.ParamInfo{},
		NumInputs:  inputIndex,
	}
	ifcWithInput := &unit.Interface{
		Parameters: map[string]*unit.ParamInfo{},
		NumInputs:  inputIndex + 1,
	}

	// empty map
	buf, err := toInput.ConnectedBuffer(wrappers)

	assert.Nil(t, buf)
	if !assert.NotNil(t, err) {
		return
	}

	// add wrapper with target ID, but is missing the param
	wrappers[id] = connect.NewUnitWrapper(nil, ifcWithoutInput, 10)

	buf, err = toInput.ConnectedBuffer(wrappers)

	assert.Nil(t, buf)
	if !assert.NotNil(t, err) {
		return
	}

	// change wrapper to one with target ID and the param
	wrappers[id] = connect.NewUnitWrapper(nil, ifcWithInput, 10)

	buf, err = toInput.ConnectedBuffer(wrappers)

	assert.NotNil(t, buf)
	assert.Nil(t, err)
}
