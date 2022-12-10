package network

import (
	"fmt"
)

type ErrUnusedEndpoint struct {
	Type    string
	Address *Address
}

func NewErrUnusedInput(addr *Address) *ErrUnusedEndpoint {
	return NewErrUnusedEndpoint("input", addr)
}

func NewErrUnusedControl(addr *Address) *ErrUnusedEndpoint {
	return NewErrUnusedEndpoint("control", addr)
}

func NewErrUnusedOutput(addr *Address) *ErrUnusedEndpoint {
	return NewErrUnusedEndpoint("output", addr)
}

func NewErrUnusedEndpoint(typ string, addr *Address) *ErrUnusedEndpoint {
	return &ErrUnusedEndpoint{
		Type:    typ,
		Address: addr,
	}
}

func (err *ErrUnusedEndpoint) Error() string {
	return fmt.Sprintf("%s %s is not used", err.Type, err.Address)
}
