package node

type Interface struct {
	Parameters map[string]*ParamInfo `json:"parameters"`
	Inputs     []string              `json:"inputs"`
}
