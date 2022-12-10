package network

import (
	"fmt"

	"github.com/dominikbraun/graph"
)

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
