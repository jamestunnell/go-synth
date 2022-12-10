package network

import "fmt"

type ErrOverusedEndpoint struct {
	Type    string
	Address *Address
}

func NewErrOverusedEndpoint(typ string, addr *Address) *ErrOverusedEndpoint {
	return &ErrOverusedEndpoint{
		Type:    typ,
		Address: addr,
	}
}

func (err *ErrOverusedEndpoint) Error() string {
	return fmt.Sprintf("%s %s is used more than once", err.Type, err.Address)
}
