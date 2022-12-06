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

func (n *Network) MarshalJSON() ([]byte, error) {
	blocks := map[string]*blockStore{}

	for name, b := range n.Blocks {
		ifc := synth.GetInterface(b)
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

		ifc := synth.GetInterface(block)

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
