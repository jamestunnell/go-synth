package param

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
)

// Type is the param type
type Type int

// Param is used to store a param value and type
type Param struct {
	typeStr string
	value   interface{}
}

// Map maps param names to values
type Map = map[string]Param

const (
	// KeyValue is the JSON key used for param value
	KeyValue = "value"
	// KeyType is the JSON key used for param type
	KeyType = "type"
	// Int indicates int64 param value
	Int = "int"
	// IntArray indicates []int64 param value
	IntArray = "[]int"
	// Float indicates float64 param value
	Float = "float"
	// FloatArray indicates []float64 param value
	FloatArray = "[]float"
	// String indicates string param value
	String = "string"
	// StringArray indicates []string param value
	StringArray = "[]string"
	// Bool indicates bool param value
	Bool = "bool"
	// BoolArray indicates []bool param value
	BoolArray = "[]bool"
)

// NewInt makes a new param with int64 value.
func NewInt(val int64) *Param {
	return &Param{typeStr: Int, value: val}
}

// NewIntArray makes a new param with []int64 value.
func NewIntArray(val []int64) *Param {
	return &Param{typeStr: IntArray, value: val}
}

// NewFloat makes a new param with float64 value.
func NewFloat(val float64) *Param {
	return &Param{typeStr: Float, value: val}
}

// NewFloatArray makes a new param with []float64 value.
func NewFloatArray(val []float64) *Param {
	return &Param{typeStr: FloatArray, value: val}
}

// NewString makes a new param with string value.
func NewString(val string) *Param {
	return &Param{typeStr: String, value: val}
}

// NewStringArray makes a new param with []string value.
func NewStringArray(val []string) *Param {
	return &Param{typeStr: StringArray, value: val}
}

// NewBool makes a new param with bool value.
func NewBool(val bool) *Param {
	return &Param{typeStr: Bool, value: val}
}

// NewBoolArray makes a new param with []bool value.
func NewBoolArray(val []bool) *Param {
	return &Param{typeStr: BoolArray, value: val}
}

// Type returns the param type
func (p *Param) Type() string {
	return p.typeStr
}

// Value returns the param value
func (p *Param) Value() interface{} {
	return p.value
}

// MarshalJSON generates param JSON data.
// Returns a non-nil error in case of failure.
func (p *Param) MarshalJSON() ([]byte, error) {
	// Use an anonymous struct which has public fields
	// (needed for json.Marshal to work)
	p2 := struct {
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	}{Type: p.typeStr, Value: p.value}

	return json.Marshal(p2)
}

// UnmarshalJSON restores a param from the given JSON data.
// Returns a non-nil error in case of failure.
func (p *Param) UnmarshalJSON(d []byte) error {
	t, err := jsonparser.GetString(d, KeyType)
	if err != nil {
		return fmt.Errorf("failed to get type: %v", err)
	}

	p.typeStr = t

	switch t {
	case Int:
		val, err := jsonparser.GetInt(d, KeyValue)
		if err != nil {
			return makeGetValErr(Int, err)
		}

		p.value = val
	case IntArray:
		var vals []int64

		err := getArrayVal(d, IntArray, func(vData []byte) error {
			return json.Unmarshal(vData, &vals)
		})

		if err != nil {
			return err
		}

		p.value = vals
	case Float:
		val, err := jsonparser.GetFloat(d, KeyValue)
		if err != nil {
			return makeGetValErr(Float, err)
		}

		p.value = val
	case FloatArray:
		var vals []float64

		err := getArrayVal(d, FloatArray, func(vData []byte) error {
			return json.Unmarshal(vData, &vals)
		})

		if err != nil {
			return err
		}

		p.value = vals
	case String:
		val, err := jsonparser.GetString(d, KeyValue)
		if err != nil {
			return makeGetValErr(String, err)
		}

		p.value = val
	case StringArray:
		var vals []string

		err := getArrayVal(d, StringArray, func(vData []byte) error {
			return json.Unmarshal(vData, &vals)
		})

		if err != nil {
			return err
		}

		p.value = vals
	case Bool:
		val, err := jsonparser.GetBoolean(d, KeyValue)
		if err != nil {
			return makeGetValErr(Bool, err)
		}

		p.value = val
	case BoolArray:
		var vals []bool

		err := getArrayVal(d, BoolArray, func(vData []byte) error {
			return json.Unmarshal(vData, &vals)
		})

		if err != nil {
			return err
		}

		p.value = vals
	default:
		return fmt.Errorf("unknown type %s", t)
	}

	return nil
}

func makeGetValErr(typStr string, err error) error {
	return fmt.Errorf("failed to get %s value: %v", typStr, err)
}

func makeParseValErr(typStr string, err error) error {
	return fmt.Errorf("failed to parse %s value: %v", typStr, err)
}

func getArrayVal(d []byte, t string, unmarshal func([]byte) error) error {
	vData, vType, _, err := jsonparser.Get(d, KeyValue)
	if err != nil {
		return makeGetValErr(t, err)
	}

	if vType != jsonparser.Array {
		return fmt.Errorf("value %s is not array type", string(vData))
	}

	err = unmarshal(vData)
	if err != nil {
		return makeParseValErr(t, err)
	}

	return nil
}
