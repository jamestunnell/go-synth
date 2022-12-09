package net

import (
	"encoding/json"
	"fmt"

	"github.com/jamestunnell/go-synth"
)

type Network struct {
	Blocks      map[string]synth.Block
	Connections []*Connection
}

type Connection struct {
	Source Address
	Dest   Address
}

type Address struct {
	Block, Port string
}

type blockStore struct {
	Path      string         `json:"path"`
	ParamVals map[string]any `json:"paramVals"`
}

type networkStore struct {
	Blocks      map[string]*blockStore `json:"blocks"`
	Connections []*Connection          `json:"connections"`
}

func NewNetwork() *Network {
	return &Network{
		Blocks:      map[string]synth.Block{},
		Connections: []*Connection{},
	}
}

func (a Address) String() string {
	return fmt.Sprintf("%s.%s", a.Block, a.Port)
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
		block, found := synth.WorkingRegistry().GetBlock(b.Path)
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

func (n *Network) CheckConnections() error {
	sourceDest := map[string]string{}
	destSource := map[string]string{}

	for _, conn := range n.Connections {
		src, dest, err := n.findConnectionEndpoints(conn)
		if err != nil {
			return fmt.Errorf("failed to find connection endpoints: %w", err)
		}

		if src.Type() != dest.Type() {
			return fmt.Errorf("source type %s does not match dest type %s", src.Type(), dest.Type())
		}

		sourceAddr := conn.Source.String()
		destAddr := conn.Dest.String()

		// check for a duplicate connection either from source or to dest
		if _, found := sourceDest[sourceAddr]; found {
			return fmt.Errorf("%s used as source address more than once", sourceAddr)
		}

		if _, found := destSource[destAddr]; found {
			return fmt.Errorf("%s used as dest address more than once", destAddr)
		}

		sourceDest[sourceAddr] = destAddr
		destSource[destAddr] = sourceAddr
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
		return nil, NewErrNotFound("output", conn.Source.Port, "source block")
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
		return nil, NewErrNotFound("input", conn.Dest.Port, "dest block")
	}

	return in, nil
}
