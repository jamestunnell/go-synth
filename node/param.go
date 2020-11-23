package node

type ParamType int

type ParamMap = map[string]interface{}

const (
	ParamTypeUInt ParamType = iota
	ParamTypeUIntArray
	ParamTypeInt
	ParamTypeIntArray
	ParamTypeFloat
	ParamTypeFloatArray
	ParamTypeString
	ParamTypeStringArray
	ParamTypeBool
	ParamTypeBoolArray
)

func (t ParamType) CheckValue(v interface{}) bool {
	switch v.(type) {
	case uint:
		return t == ParamTypeUInt
	case []uint:
		return t == ParamTypeUIntArray
	case int:
		return t == ParamTypeInt
	case []int:
		return t == ParamTypeIntArray
	case float64:
		return t == ParamTypeFloat
	case []float64:
		return t == ParamTypeFloatArray
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
	case ParamTypeUInt:
		return "UInt"
	case ParamTypeUIntArray:
		return "UIntArray"
	case ParamTypeInt:
		return "Int"
	case ParamTypeIntArray:
		return "IntArray"
	case ParamTypeFloat:
		return "Float"
	case ParamTypeFloatArray:
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
