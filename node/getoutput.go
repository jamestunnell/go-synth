package node

import "fmt"

func GetOutput(nmap map[string]*Node, name string) *Buffer {
	node, found := nmap[name]
	if !found {
		panic(fmt.Errorf("failed to find node %s", name))
	}

	return node.Output
}
