package param

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/buger/jsonparser"
)

// Type is the param type
type Type = string

// Param is used to store a param value and type
type Param struct {
	typ   Type
	value interface{}
}

// Map maps param names to values
type Map = map[string]*Param

const (
	// KeyValue is the JSON key used for param value
	KeyValue = Type("value")
	// KeyType is the JSON key used for param type
	KeyType = Type("type")
	// Int indicates int64 param value
	Int = Type("int")
	// IntArray indicates []int64 param value
	IntArray = Type("[]int")
	// Float indicates float64 param value
	Float = Type("float")
	// FloatArray indicates []float64 param value
	FloatArray = Type("[]float")
	// String indicates string param value
	String = Type("string")
	// StringArray indicates []string param value
	StringArray = Type("[]string")
	// Bool indicates bool param value
	Bool = Type("bool")
	// BoolArray indicates []bool param value
	BoolArray = Type("[]bool")
)

// NewInt makes a new param with int64 value.
func NewInt(val int64) *Param {
	return &Param{typ: Int, value: val}
}

// NewIntArray makes a new param with []int64 value.
func NewIntArray(val []int64) *Param {
	return &Param{typ: IntArray, value: val}
}

// NewFloat makes a new param with float64 value.
func NewFloat(val float64) *Param {
	return &Param{typ: Float, value: val}
}

// NewFloatArray makes a new param with []float64 value.
func NewFloatArray(val []float64) *Param {
	return &Param{typ: FloatArray, value: val}
}

// NewString makes a new param with string value.
func NewString(val string) *Param {
	return &Param{typ: String, value: val}
}

// NewStringArray makes a new param with []string value.
func NewStringArray(val []string) *Param {
	return &Param{typ: StringArray, value: val}
}

// NewBool makes a new param with bool value.
func NewBool(val bool) *Param {
	return &Param{typ: Bool, value: val}
}

// NewBoolArray makes a new param with []bool value.
func NewBoolArray(val []bool) *Param {
	return &Param{typ: BoolArray, value: val}
}

// Type returns the param type
func (p *Param) Type() Type {
	return p.typ
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
		Type  Type        `json:"type"`
		Value interface{} `json:"value"`
	}{Type: p.typ, Value: p.value}

	return json.Marshal(p2)
}

// UnmarshalJSON restores a param from the given JSON data.
// Returns a non-nil error in case of failure.
func (p *Param) UnmarshalJSON(d []byte) error {
	t, err := jsonparser.GetString(d, KeyType)
	if err != nil {
		return fmt.Errorf("failed to get type: %v", err)
	}

	p.typ = t

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

func makeGetValErr(typ Type, err error) error {
	return fmt.Errorf("failed to get %s value: %v", typ, err)
}

func makeParseValErr(typ Type, err error) error {
	return fmt.Errorf("failed to parse %s value: %v", typ, err)
}

func getArrayVal(d []byte, t string, unmarshal func([]byte) error) error {
	vData, vType, _, err := jsonparser.Get(d, KeyValue)
	if err != nil {
		return makeGetValErr(t, err)
	}

	if vType != jsonparser.Array {
		return errors.New("value is not array type")
	}

	err = unmarshal(vData)
	if err != nil {
		return makeParseValErr(t, err)
	}

	return nil
}
