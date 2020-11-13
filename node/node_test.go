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

func TestNodeNewNodeMissingInput(t *testing.T) {
	defer func() { recover() }()

	node.NewNode(&nodetest.MulAdd{}, node.Map{}, node.Map{})

	t.Error("should not be reached")
}

func TestNodeNewNodeMissingControl(t *testing.T) {
	in := node.NewConst(0.0)
	mulk := node.NewConst(5.0)
	addk := node.NewConst(2.5)

	n1 := node.NewNode(
		&nodetest.MulAdd{},
		node.Map{"In": in},
		node.Map{"MulK": mulk})
	addk2, found := n1.Controls["AddK"]

	if assert.True(t, found) {
		assert.Equal(t, nodetest.MulAddDefaultAddK, addk2.Core.(*node.Const).Value)
	}

	n2 := node.NewNode(
		&nodetest.MulAdd{},
		node.Map{"In": in},
		node.Map{"AddK": addk})
	mulk2, found := n2.Controls["MulK"]

	if assert.True(t, found) {
		assert.Equal(t, nodetest.MulAddDefaultMulK, mulk2.Core.(*node.Const).Value)
	}
}

func TestNodeUnmarshalHappyPath(t *testing.T) {
	_, d := marshaledNodeRegisterCore(t)

	var n2 node.Node

	err := json.Unmarshal(d, &n2)

	assert.Nil(t, err)

	// Should still unmarshal fine after removing a control from JSON
	d = jsonparser.Delete(d, "controls", "MulK")

	err = json.Unmarshal(d, &n2)

	assert.Nil(t, err)

	if assert.Contains(t, n2.Controls, "MulK") {
		mulk := n2.Controls["MulK"]

		assert.Equal(t, mulk.Core.(*node.Const).Value, nodetest.MulAddDefaultMulK)
	}
}

func TestNodeUnmarshalCoreNotInRegistry(t *testing.T) {
	c, d := marshaledNodeRegisterCore(t)

	node.WorkingRegistry().UnregisterCore(c)

	var n2 node.Node

	err := json.Unmarshal(d, &n2)

	assert.NotNil(t, err)
}

func TestNodeUnmarshalMissingInput(t *testing.T) {
	_, d := marshaledNodeRegisterCore(t)
	s := string(d)

	d = []byte(strings.Replace(s, "In", "Ex", 1))

	var n2 node.Node

	defer func() { recover() }()

	json.Unmarshal(d, &n2)

	t.Error("should not be reached")
}

func marshaledNodeRegisterCore(t *testing.T) (node.Core, []byte) {
	in := node.NewConst(0.0)
	mulk := node.NewConst(2.5)
	addk := node.NewConst(7.7)
	c := &nodetest.MulAdd{}
	n := node.NewNode(c, node.Map{"In": in}, node.Map{"MulK": mulk, "AddK": addk})

	node.WorkingRegistry().RegisterCore(c)

	d, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}

	return c, d
}
