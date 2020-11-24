package param_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/util/param"
)

func TestParamUnmarshalMissingType(t *testing.T) {
	str := fmt.Sprintf(`{"%s":5}`, param.KeyValue)
	testUnmarshalFail(t, str)
}

func TestIntParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewInt(5), param.Int)
}

func TestIntParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewInt(5))
}

func TestIntParamUnmarshalBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewInt(5), "5.0")
}

func TestIntArrayParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewIntArray([]int64{5, 6, 7}), param.IntArray)
}

func TestIntArrayParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewIntArray([]int64{5}))
}

func TestIntArrayParamUnmarshalNotAnArray(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewIntArray([]int64{5}), "5")
}

func TestIntArrayParamUnmarshalArrayWithBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewIntArray([]int64{5}), `[5, "abc"]`)
}

func TestFloatParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewFloat(5.5), param.Float)
}

func TestFloatParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewFloat(5.5))
}

func TestFloatParamUnmarshalBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewFloat(5.5), `"abc"`)
}

func TestFloatArrayParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewFloatArray([]float64{5.5}), param.FloatArray)
}

func TestFloatArrayParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewFloatArray([]float64{5.2}))
}

func TestFloatArrayParamUnmarshalNotAnArray(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewFloatArray([]float64{5.2}), "5.2")
}

func TestFloatArrayParamUnmarshalArrayWithBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewFloatArray([]float64{5}), `[5.2, "abc"]`)
}

func TestStringParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewString("abc"), param.String)
}

func TestStringParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewString("abc"))
}

func TestStringParamUnmarshalBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewString("abc"), "5")
}

func TestStringArrayParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewStringArray([]string{"abc"}), param.StringArray)
}

func TestStringArrayParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewStringArray([]string{"abc"}))
}

func TestStringArrayParamUnmarshalNotAnArray(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewStringArray([]string{"abc"}), `"abc"`)
}

func TestStringArrayParamUnmarshalArrayWithBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewStringArray([]string{"abc"}), `["abc", 5]`)
}

func TestBoolParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewBool(true), param.Bool)
}

func TestBoolParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewBool(true))
}

func TestBoolParamUnmarshalBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewBool(true), "5")
}

func TestBoolArrayParamHappyPath(t *testing.T) {
	testParamHappyPath(t, param.NewBoolArray([]bool{true}), param.BoolArray)
}

func TestBoolArrayParamUnmarshalMissingVal(t *testing.T) {
	testParamUnmarshalMissingVal(t, param.NewBoolArray([]bool{true}))
}

func TestBoolArrayParamUnmarshalNotAnArray(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewBoolArray([]bool{true}), `true`)
}

func TestBoolArrayParamUnmarshalArrayWithBadVal(t *testing.T) {
	testParamUnmarshalBadVal(t, param.NewBoolArray([]bool{true}), `[true, 7]`)
}

func testParamHappyPath(t *testing.T, p *param.Param, expectedType string) {
	if !assert.Equal(t, expectedType, p.Type()) {
		return
	}

	d, err := json.Marshal(p)
	if !assert.NoError(t, err) {
		return
	}

	var p2 param.Param
	err = json.Unmarshal(d, &p2)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, p.Value(), p2.Value())
	assert.Equal(t, p.Type(), p2.Type())
}

func testParamUnmarshalMissingVal(t *testing.T, p *param.Param) {
	str := fmt.Sprintf(`{"%s":"%s"}`, param.KeyType, p.Type())
	testUnmarshalFail(t, str)
}

func testParamUnmarshalBadVal(t *testing.T, p *param.Param, badValStr string) {
	str := fmt.Sprintf(
		`{"%s":"%s","%s":%s}`, param.KeyType, p.Type(), param.KeyValue, badValStr)
	testUnmarshalFail(t, str)
}

func testUnmarshalFail(t *testing.T, jsonStr string) {
	var p2 param.Param

	err := json.Unmarshal([]byte(jsonStr), &p2)
	assert.Error(t, err)
}
