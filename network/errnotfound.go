package network

import "fmt"

type ErrNotFound struct {
	Type, Name, Place string
}

func NewErrNotFound(typ, name, place string) *ErrNotFound {
	return &ErrNotFound{
		Type:  typ,
		Name:  name,
		Place: place,
	}
}

func (err *ErrNotFound) Error() string {
	const strFmt = "%s %s not found in %s"

	return fmt.Sprintf(strFmt, err.Type, err.Name, err.Place)
}
