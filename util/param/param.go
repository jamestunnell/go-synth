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
	// Ints indicates []int64 param value
	Ints = Type("[]int")
	// Float indicates float64 param value
	Float = Type("float")
	// Floats indicates []float64 param value
	Floats = Type("[]float")
	// String indicates string param value
	String = Type("string")
	// Strings indicates []string param value
	Strings = Type("[]string")
	// Bool indicates bool param value
	Bool = Type("bool")
	// Bools indicates []bool param value
	Bools = Type("[]bool")
)

// NewInt makes a new param with int64 value.
func NewInt(val int64) *Param {
	return &Param{typ: Int, value: val}
}

// NewInts makes a new param with []int64 value.
func NewInts(val []int64) *Param {
	return &Param{typ: Ints, value: val}
}

// NewFloat makes a new param with float64 value.
func NewFloat(val float64) *Param {
	return &Param{typ: Float, value: val}
}

// NewFloats makes a new param with []float64 value.
func NewFloats(val []float64) *Param {
	return &Param{typ: Floats, value: val}
}

// NewString makes a new param with string value.
func NewString(val string) *Param {
	return &Param{typ: String, value: val}
}

// NewStrings makes a new param with []string value.
func NewStrings(val []string) *Param {
	return &Param{typ: Strings, value: val}
}

// NewBool makes a new param with bool value.
func NewBool(val bool) *Param {
	return &Param{typ: Bool, value: val}
}

// NewBools makes a new param with []bool value.
func NewBools(val []bool) *Param {
	return &Param{typ: Bools, value: val}
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
	case Ints:
		var vals []int64

		err := getArrayVal(d, Ints, func(vData []byte) error {
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
	case Floats:
		var vals []float64

		err := getArrayVal(d, Floats, func(vData []byte) error {
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
	case Strings:
		var vals []string

		err := getArrayVal(d, Strings, func(vData []byte) error {
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
	case Bools:
		var vals []bool

		err := getArrayVal(d, Bools, func(vData []byte) error {
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
