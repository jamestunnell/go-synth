package network

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"

	"github.com/dominikbraun/graph"
	"github.com/jamestunnell/go-synth"
)

type Network struct {
	Blocks      map[string]synth.Block
	Connections Connections
}

type blockStore struct {
	Path      string         `json:"path"`
	ParamVals map[string]any `json:"paramVals"`
}

type networkStore struct {
	Blocks      map[string]*blockStore `json:"blocks"`
	Connections []*Connection          `json:"connections"`
}

func New() *Network {
	return &Network{
		Blocks:      map[string]synth.Block{},
		Connections: Connections{},
	}
}

func (n *Network) Equal(other *Network) bool {
	if len(n.Blocks) != len(other.Blocks) {
		return false
	}

	for name, b := range n.Blocks {
		b2, found := other.Blocks[name]
		if !found {
			return false
		}

		if synth.BlockPath(b) != synth.BlockPath(b2) {
			return false
		}

		pVals := synth.BlockInterface(b).ParamVals()
		pVals2 := synth.BlockInterface(b).ParamVals()
		for name, val := range pVals {
			val2, found := pVals2[name]
			if !found {
				return false
			}

			if !reflect.DeepEqual(val, val2) {
				return false
			}
		}
	}

	return n.Connections.Equal(other.Connections)
}

func (n *Network) BlockOrder() ([]string, error) {
	g := graph.New(graph.StringHash, graph.Directed(), graph.Acyclic())

	for name := range n.Blocks {
		g.AddVertex(name)
	}

	for _, conn := range n.Connections {
		g.AddEdge(conn.Source.Block, conn.Dest.Block)
	}

	order, err := graph.TopologicalSort(g)
	if err != nil {
		return []string{}, fmt.Errorf("topological graph sort failed: %w", err)
	}

	return order, nil
}

func (n *Network) MarshalJSON() ([]byte, error) {
	blocks := map[string]*blockStore{}

	for name, b := range n.Blocks {
		ifc := synth.BlockInterface(b)
		path := synth.BlockPath(b)
		paramVals := ifc.ParamVals()

		blocks[name] = &blockStore{
			Path:      path,
			ParamVals: paramVals,
		}
	}

	ns := &networkStore{
		Blocks:      blocks,
		Connections: n.Connections,
	}

	return json.Marshal(ns)
}

func (n *Network) UnmarshalJSON(d []byte) error {
	var ns networkStore

	if err := json.Unmarshal(d, &ns); err != nil {
		return fmt.Errorf("failed to unmarshal JSON as networkStore: %w", err)
	}

	blocks := map[string]synth.Block{}

	for name, b := range ns.Blocks {
		block, found := synth.WorkingRegistry().MakeBlock(b.Path)
		if !found {
			return NewErrNotFound("path", b.Path, "block registry")
		}

		ifc := synth.BlockInterface(block)

		for name, val := range b.ParamVals {
			param, found := ifc.Params[name]
			if !found {
				return NewErrNotFound("param", name, "block")
			}

			if err := param.SetValue(val); err != nil {
				return fmt.Errorf("failed to set param val: %w", err)
			}
		}

		blocks[name] = block
	}

	n.Blocks = blocks
	n.Connections = ns.Connections

	return nil
}

func (n *Network) AddDefaultBlocks() error {
	defaultBlocks := map[string]synth.Block{}
	defaultConns := Connections{}

	for bName, b := range n.Blocks {
		ifc := synth.BlockInterface(b)

		for cName, c := range ifc.Controls {
			dest := NewAddress(bName, cName)

			if _, found := n.Connections.FindByDest(dest); found {
				continue
			}

			dblk, err := synth.NewConstFromAny(c.DefaultVal())
			if err != nil {
				return fmt.Errorf("failed to make const block for val %v", c.DefaultVal())
			}

			dblkName := fmt.Sprintf("Default_%s_%d", cName, rand.Intn(5000))

			defaultBlocks[dblkName] = dblk

			dconn := NewConnection(NewAddress(dblkName, "Out"), dest)

			defaultConns = append(defaultConns, dconn)
		}

	}

	for name, block := range defaultBlocks {
		n.Blocks[name] = block
	}

	n.Connections = append(n.Connections, defaultConns...)

	return nil
}

func (n *Network) Validate() []error {
	errs := []error{}

	for _, conn := range n.Connections {
		if err := n.checkConnection(conn); err != nil {
			err = fmt.Errorf("connection %s is not valid: %w", conn, err)

			errs = append(errs, err)
		}
	}

	for name, blk := range n.Blocks {
		if err := n.checkBlock(blk, name); err != nil {
			err = fmt.Errorf("block %s is not valid: %w", name, err)

			errs = append(errs, err)
		}
	}

	for _, src := range n.Connections.OverusedSources() {
		errs = append(errs, NewErrOverusedEndpoint("source", src))
	}

	for _, dest := range n.Connections.OverusedDests() {
		errs = append(errs, NewErrOverusedEndpoint("dest", dest))
	}

	return errs
}

// checkConnection ensures that connection endpoints can be found in
// the network blocks and the endpoint types match.
func (n *Network) checkConnection(conn *Connection) error {
	src, dest, err := n.findConnectionEndpoints(conn)
	if err != nil {
		return fmt.Errorf("failed to find connection endpoints: %w", err)
	}

	if src.Type() != dest.Type() {
		return fmt.Errorf("source type %s does not match dest type %s", src.Type(), dest.Type())
	}

	return nil
}

func (n *Network) checkBlock(blk synth.Block, name string) error {
	ifc := synth.BlockInterface(blk)

	untargeted := []string{}

	for inName := range ifc.Inputs {
		dest := &Address{Block: name, Port: inName}

		_, found := n.Connections.FindByDest(dest)
		if !found {
			untargeted = append(untargeted, inName)
		}
	}

	if len(untargeted) > 0 {
		return NewErrUntargetedInputs(untargeted)
	}

	return nil
}

func (n *Network) findConnectionEndpoints(conn *Connection) (synth.Output, synth.Input, error) {
	out, err := n.findSource(conn)
	if err != nil {
		return nil, nil, err
	}

	in, err := n.findDest(conn)
	if err != nil {
		return nil, nil, err
	}

	return out, in, nil
}

func (n *Network) findSource(conn *Connection) (synth.Output, error) {
	b, found := n.Blocks[conn.Source.Block]
	if !found {
		return nil, NewErrNotFound("source block", conn.Source.Block, "network")
	}

	ifc := synth.BlockInterface(b)

	out, found := ifc.Outputs[conn.Source.Port]
	if !found {
		return nil, NewErrNotFound("output", conn.Source.Port, "source block "+conn.Source.Block)
	}

	return out, nil
}

func (n *Network) findDest(conn *Connection) (synth.Input, error) {
	b, found := n.Blocks[conn.Dest.Block]
	if !found {
		return nil, NewErrNotFound("dest block", conn.Dest.Block, "network")
	}

	ifc := synth.BlockInterface(b)

	in, found := ifc.Inputs[conn.Dest.Port]
	if !found {
		ctrl, found := ifc.Controls[conn.Dest.Port]
		if !found {
			return nil, NewErrNotFound("input/control", conn.Dest.Port, "dest block "+conn.Dest.Block)
		}

		return ctrl, nil
	}

	return in, nil
}
