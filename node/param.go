package node

type ParamType int

type ParamMap = map[string]interface{}

const (
	ParamTypeNumber ParamType = iota
	ParamTypeNumberArray
	ParamTypeString
	ParamTypeStringArray
	ParamTypeBool
	ParamTypeBoolArray
)

func (t ParamType) CheckValue(v interface{}) bool {
	switch v.(type) {
	case float64:
		return t == ParamTypeNumber
	case []float64:
		return t == ParamTypeNumberArray
	case string:
		return t == ParamTypeString
	case []string:
		return t == ParamTypeStringArray
	case bool:
		return t == ParamTypeBool
	case []bool:
		return t == ParamTypeBoolArray
	}

	return false
}

func (t ParamType) String() string {
	switch t {
	case ParamTypeNumber:
		return "Float"
	case ParamTypeNumberArray:
		return "Array"
	case ParamTypeString:
		return "String"
	case ParamTypeStringArray:
		return "StringArray"
	case ParamTypeBool:
		return "Bool"
	case ParamTypeBoolArray:
		return "BoolArray"
	}

	return ""
}
