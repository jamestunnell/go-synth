package network

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Address struct {
	Block, Port string
}

func NewAddress(block, port string) *Address {
	return &Address{
		Block: block,
		Port:  port,
	}
}

func (a *Address) String() string {
	return fmt.Sprintf("%s.%s", a.Block, a.Port)
}

func (a *Address) MarshalJSON() ([]byte, error) {
	str := a.String()

	return json.Marshal(str)
}

func (a *Address) UnmarshalJSON(d []byte) error {
	var str string

	if err := json.Unmarshal(d, &str); err != nil {
		return fmt.Errorf("failed to unmarshal JSON as string: %w", err)
	}

	return a.Parse(str)
}

func (a *Address) Equal(other *Address) bool {
	return a.Block == other.Block && a.Port == other.Port
}

func (a *Address) Parse(addr string) error {
	strs := strings.Split(addr, ".")
	if len(strs) != 2 {
		return fmt.Errorf("failed to split address '%s' into two substrings", addr)
	}

	a.Block = strs[0]
	a.Port = strs[1]

	return nil
}
