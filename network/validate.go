package network

import (
	"fmt"

	"github.com/jamestunnell/go-synth"
)

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
