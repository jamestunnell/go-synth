package constraints

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jamestunnell/go-synth/pkg/unit"
)

func UnmarshalConstraintJSON(d []byte) (unit.Constraint, error) {
	var obj map[string]interface{}

	err := json.Unmarshal(d, &obj)
	if err != nil {
		return nil, err
	}

	_, found := obj["value"]
	if found {
		var sv SingleValue
		err := json.Unmarshal(d, &sv)
		if err != nil {
			return nil, err
		}

		switch sv.Type {
		case typeStrLess:
			return &Less{sv}, nil
		case typeStrLessEqual:
			return &LessEqual{sv}, nil
		case typeStrGreater:
			return &Greater{sv}, nil
		case typeStrGreaterEqual:
			return &GreaterEqual{sv}, nil
		case typeStrEqual:
			return &Equal{sv}, nil
		default:
			return nil, fmt.Errorf("unknown single-value constraint type %s", sv.Type)
		}
	}

	_, found = obj["values"]

	if found {
		var mv MultiValue
		err := json.Unmarshal(d, &mv)
		if err != nil {
			return nil, err
		}

		switch mv.Type {
		case typeStrInSet:
			return &InSet{mv}, nil
		case typeStrNotInSet:
			return &NotInSet{mv}, nil
		default:
			return nil, fmt.Errorf("unknown multi-value constraint type %s", mv.Type)
		}
	}

	return nil, errors.New("invalid format for constraint JSON")
}
