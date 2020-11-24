package node_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/nodetest"
	"github.com/stretchr/testify/assert"
)

func TestNodeInitializeMissingInput(t *testing.T) {
	n := testNode()

	delete(n.Inputs, nodetest.InputName)

	assert.Error(t, n.Initialize(100.0, 4))
}

func TestNodeInitializeMissingParam(t *testing.T) {
	n := testNode()

	delete(n.Params, nodetest.ParamName)

	assert.Error(t, n.Initialize(100.0, 4))
}

func TestNodeInitializeBadParamVal(t *testing.T) {
	n := testNode()

	n.Params[nodetest.ParamName] = nodetest.BadParamVal

	assert.Error(t, n.Initialize(100.0, 4))
}

func TestNodeInitializeMissingControl(t *testing.T) {
	n := testNode()

	delete(n.Controls, nodetest.ControlName)

	assert.NoError(t, n.Initialize(100.0, 4))

	if assert.Contains(t, n.Controls, nodetest.ControlName) {
		assert.Equal(t, nodetest.ControlDefault, n.Controls[nodetest.ControlName].Core.(*node.Const).Value)
	}
}

func TestNodeUnmarshalHappyPath(t *testing.T) {
	c, d := marshaledNode(t)

	node.WorkingRegistry().Register(c)

	var n2 node.Node

	err := json.Unmarshal(d, &n2)

	assert.Nil(t, err)

	// Should still unmarshal fine after removing a control from JSON
	d = jsonparser.Delete(d, "controls", nodetest.ControlName)

	if !assert.NoError(t, json.Unmarshal(d, &n2)) {
		return
	}

	if !assert.NoError(t, n2.Initialize(100.0, 3)) {
		return
	}

	if assert.Contains(t, n2.Controls, nodetest.ControlName) {
		ctrl := n2.Controls[nodetest.ControlName]

		assert.Equal(t, nodetest.ControlDefault, ctrl.Core.(*node.Const).Value)
	}
}

func TestNodeUnmarshalCoreNotInRegistry(t *testing.T) {
	c, d := marshaledNode(t)

	node.WorkingRegistry().Unregister(node.CorePath(c))

	var n2 node.Node

	assert.Error(t, json.Unmarshal(d, &n2))
}

func TestNodeUnmarshalMissingInput(t *testing.T) {
	c, d := marshaledNode(t)
	s := string(d)

	node.WorkingRegistry().Register(c)

	d = []byte(strings.Replace(s, nodetest.InputName, "Ex", 1))

	var n2 node.Node

	assert.Error(t, json.Unmarshal(d, &n2))
}

func marshaledNode(t *testing.T) (node.Core, []byte) {
	n := testNode()

	d, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("node JSON: %s\n", string(d))

	return n.Core, d
}

func testNode() *node.Node {
	return &node.Node{
		Core:     &nodetest.TestCore{},
		Inputs:   node.Map{nodetest.InputName: node.NewConst(227)},
		Controls: node.Map{nodetest.ControlName: node.NewConst(54)},
		Params:   node.ParamMap{nodetest.ParamName: 2.0},
	}
}
