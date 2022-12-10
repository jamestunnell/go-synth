package network_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/network"
	"github.com/jamestunnell/go-synth/unit/gen"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
	"github.com/jamestunnell/go-synth/unit/proc"
	"github.com/jamestunnell/go-synth/unit/proc/math"
)

func TestNetworkHappyPath(t *testing.T) {
	n := network.New()

	n.Blocks["A"] = osc.NewTriangle()
	n.Blocks["B"] = osc.NewSine()
	n.Blocks["C"] = math.NewMul()
	n.Blocks["D"] = synth.NewConst(20.0)
	n.Blocks["E"] = synth.NewConst(200.0)
	n.Blocks["F"] = synth.NewMonoTerminal()

	n.Connections = network.Connections{
		mustParseConn(t, "D.Out -> A.Freq"),
		mustParseConn(t, "E.Out -> B.Freq"),
		mustParseConn(t, "A.Out -> C.In1"),
		mustParseConn(t, "B.Out -> C.In2"),
		mustParseConn(t, "C.Out -> F.In"),
	}

	testNetworkHappyPath(t, n)
}

// func TestNetworkWithIslands(t *testing.T) {
// 	n := network.New()

// 	n.Blocks["A"] = osc.NewTriangle()
// 	n.Blocks["B"] = osc.NewSine()
// 	n.Blocks["D"] = synth.NewConst(20.0)
// 	n.Blocks["E"] = synth.NewConst(200.0)

// 	n.Connections = network.Connections{
// 		mustParseConn(t, "D.Out -> A.Freq"),
// 		mustParseConn(t, "E.Out -> B.Freq"),
// 	}

// 	testNetwork(t, n)
// }

func testNetworkHappyPath(t *testing.T, n *network.Network) {
	n.AddDefaultBlocks()

	require.Empty(t, n.Validate())

	testSerialization(t, n)

	order, err := n.BlockOrder()

	require.NoError(t, err)

	assert.Len(t, order, len(n.Blocks))
}

func testSerialization(t *testing.T, n *network.Network) {
	d, err := n.MarshalJSON()
	require.NoError(t, err)

	gen.RegisterBuiltin(synth.WorkingRegistry())
	proc.RegisterBuiltin(synth.WorkingRegistry())

	var n2 network.Network

	require.NoError(t, json.Unmarshal(d, &n2))

	assert.True(t, n.Equal(&n2))
}

func mustParseConn(t *testing.T, conn string) *network.Connection {
	strs := strings.Split(conn, "->")

	src := &network.Address{}
	dest := &network.Address{}

	require.NoError(t, src.Parse(strings.TrimSpace(strs[0])))
	require.NoError(t, dest.Parse(strings.TrimSpace(strs[1])))

	return &network.Connection{Source: src, Dest: dest}
}
