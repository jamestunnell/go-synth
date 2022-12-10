package network

import "fmt"

func (n *Network) ApplyConnections() error {
	for _, conn := range n.Connections {
		out, in, err := n.findConnectionEndpoints(conn)
		if err != nil {
			return fmt.Errorf("failed to find endpoints: %w", err)
		}

		if err = in.Connect(out); err != nil {
			return fmt.Errorf("failed to apply connection %s: %w", conn, err)
		}
	}
	return nil
}
