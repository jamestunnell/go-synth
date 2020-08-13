package connect_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jamestunnell/go-synth/connect"
	"github.com/jamestunnell/go-synth/unit"
	"github.com/stretchr/testify/assert"
)

func TestToParam(t *testing.T) {
	const paramName = "abc"

	id := uuid.New()
	wrappers := map[uuid.UUID]*connect.UnitWrapper{}
	toParam := connect.ToParam{TargetID: id, ParamName: paramName}
	ifcWithoutParam := &unit.Interface{
		Parameters: map[string]*unit.ParamInfo{"xyz": &unit.ParamInfo{}},
	}
	ifcWithParam := &unit.Interface{
		Parameters: map[string]*unit.ParamInfo{paramName: &unit.ParamInfo{}},
	}

	// empty map
	buf, err := toParam.ConnectedBuffer(wrappers)

	assert.Nil(t, buf)
	if !assert.NotNil(t, err) {
		return
	}

	// add wrapper with target ID, but is missing the param
	wrappers[id] = connect.NewUnitWrapper(nil, ifcWithoutParam, 10)

	buf, err = toParam.ConnectedBuffer(wrappers)

	assert.Nil(t, buf)
	if !assert.NotNil(t, err) {
		return
	}

	// change wrapper to one with target ID and the param
	wrappers[id] = connect.NewUnitWrapper(nil, ifcWithParam, 10)

	buf, err = toParam.ConnectedBuffer(wrappers)

	assert.NotNil(t, buf)
	assert.Nil(t, err)
}
