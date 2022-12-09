package network

import "fmt"

type Connection struct {
	Source *Address `json:"source"`
	Dest   *Address `json:"dest"`
}

func (conn *Connection) String() string {
	return fmt.Sprintf("%s -> %s", conn.Source, conn.Dest)
}
