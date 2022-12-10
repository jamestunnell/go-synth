package network

import (
	"encoding/json"
	"fmt"

	"github.com/jamestunnell/go-synth"
)

type Network struct {
	Blocks      BlockMap
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
		Blocks:      BlockMap{},
		Connections: Connections{},
	}
}

func (n *Network) Equal(other *Network) bool {
	return n.Blocks.Equal(other.Blocks) && n.Connections.Equal(other.Connections)
}

func (n *Network) InitializeBlocks(srate float64, outDepth int) error {
	for name, blk := range n.Blocks {
		if err := blk.Initialize(srate, outDepth); err != nil {
			return fmt.Errorf("failed to initialize block %s: %w", name, err)
		}
	}

	return nil
}

func (n *Network) MakeConfigureAndRunFunc() (func(), error) {
	order, err := n.BlockOrder()
	if err != nil {
		return nil, fmt.Errorf("failed to make block order: %w", err)
	}

	blocks := make([]synth.Block, len(order))
	for i := 0; i < len(order); i++ {
		blkName := order[i]

		blk, found := n.Blocks[blkName]
		if !found {
			err := fmt.Errorf("block %s not found", blkName)

			return nil, err
		}

		blocks[i] = blk
	}

	f := func() {
		for _, blk := range blocks {
			blk.Configure()
		}

		for _, blk := range blocks {
			blk.Run()
		}
	}

	return f, nil
}

func (n *Network) TerminalBlocks() []synth.Block {
	terminals := []synth.Block{}

	for _, blk := range n.Blocks {
		if synth.IsTerminal(blk) {
			terminals = append(terminals, blk)
		}
	}

	return terminals
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
