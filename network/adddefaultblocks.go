package network

import (
	"fmt"
	"math/rand"

	"github.com/jamestunnell/go-synth"
)

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
