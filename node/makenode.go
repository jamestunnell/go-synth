package node

import (
	"fmt"
	"reflect"
	"unicode"
)

func MakeNode(core Core, depth int) (*Node, error) {
	inputNodes := []*Node{}
	paramNodes := []*Node{}

	if depth < 1 {
		return nil, fmt.Errorf("buffer depth %d is not positive", depth)
	}

	v1 := reflect.ValueOf(core)
	coreType := v1.Elem().Type()
	v2 := reflect.New(coreType)

	// fmt.Printf("core type: %s, buffer depth: %d\n", v1.Elem().Type(), depth)

	ifc := core.GetInterface()

	for _, input := range ifc.Inputs {
		f1 := v1.Elem().FieldByName(input)
		f2 := v2.Elem().FieldByName(input)

		if !f1.IsValid() {
			return nil, fmt.Errorf("missing input %s", input)
		}

		if !f1.CanInterface() {
			return nil, fmt.Errorf("input %s is not an interface{}", input)
		}

		if f1.IsNil() {
			return nil, fmt.Errorf("input %s is not set", input)
		}

		inCore, ok := f1.Interface().(Core)
		if !ok {
			return nil, fmt.Errorf("input %s is not a Core", input)
		}

		// inputs should be nodes with the same depth as our output buffer
		inNode, err := MakeNode(inCore, depth)
		if err != nil {
			return nil, err
		}

		inputNodes = append(inputNodes, inNode)

		f2.Set(reflect.ValueOf(inNode))
	}

	for paramName, paramInfo := range ifc.Parameters {
		f1 := v1.Elem().FieldByName(paramName)
		f2 := v2.Elem().FieldByName(paramName)

		if !f1.IsValid() {
			return nil, fmt.Errorf("missing param %s", paramName)
		}

		if !f1.CanInterface() {
			return nil, fmt.Errorf("param %s is not an interface{}", paramName)
		}

		if f1.IsNil() {
			if paramInfo.Required {
				return nil, fmt.Errorf("required param %s is not set", paramName)
			} else {
				paramNode := MakeConstNode(paramInfo.Default)

				f2.Set(reflect.ValueOf(paramNode))

				paramNodes = append(paramNodes, paramNode)

				continue
			}
		}

		var paramNode *Node

		param := f1.Interface()
		switch v := param.(type) {
		case float64:
			paramNode = MakeConstNode(v)
		case float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			paramNode = MakeConstNode(v.(float64))
		case Core:
			var err error
			// param nodes should always have depth 1
			paramNode, err = MakeNode(v, 1)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported param value type %s", f1.Elem().Type().Name())
		}

		paramNodes = append(paramNodes, paramNode)

		f2.Set(reflect.ValueOf(paramNode))
	}

	// Other - these are public fields that are not part of the interface
	// we transfer these across as-is in case they are needed
	for i := 0; i < coreType.NumField(); i++ {
		fieldName := coreType.Field(i).Name

		// ignore unexported fields
		if !unicode.IsUpper(rune(fieldName[0])) {
			continue
		}

		_, isParam := ifc.Parameters[fieldName]

		if isParam {
			continue
		}

		isInput := false
		for _, inputName := range ifc.Inputs {
			if fieldName == inputName {
				isInput = true
				break
			}
		}

		if isInput {
			continue
		}

		f1 := v1.Elem().FieldByName(fieldName)
		f2 := v2.Elem().FieldByName(fieldName)
		f2.Set(f1)
	}

	node := &Node{
		Out:        NewBuffer(depth),
		Core:       v2.Interface().(Core),
		InputNodes: inputNodes,
		ParamNodes: paramNodes,
	}

	return node, nil
}

func MakeConstNode(val float64) *Node {
	out := NewBuffer(1)
	out.Values[0] = val

	return &Node{Out: out, Core: &NoOpCore{}}
}
