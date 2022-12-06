package synth

import "reflect"

type Interface struct {
	Inputs   map[string]Input
	Controls map[string]Control
	Params   map[string]Param
	Outputs  map[string]Output
}

func GetInterface(b Block) *Interface {
	ifc := NewInterface()

	ifc.Extract(b)

	return ifc
}

func NewInterface() *Interface {
	return &Interface{
		Inputs:   map[string]Input{},
		Controls: map[string]Control{},
		Params:   map[string]Param{},
		Outputs:  map[string]Output{},
	}
}

func (ifc *Interface) Extract(b Block) {
	st := reflect.TypeOf(b).Elem()
	sv := reflect.ValueOf(b).Elem()

	for i := 0; i < st.NumField(); i++ {
		stf := st.Field(i)

		if !stf.IsExported() {
			continue
		}

		svf := sv.Field(i)

		if !svf.CanInterface() {
			continue
		}

		f := svf.Interface()

		switch v := f.(type) {
		case Block:
			ifc.Extract(v)
		case Control:
			ifc.Controls[stf.Name] = v
		case Input:
			ifc.Inputs[stf.Name] = v
		case Param:
			ifc.Params[stf.Name] = v
		case Output:
			ifc.Outputs[stf.Name] = v
		}
	}
}

func (ifc *Interface) ParamVals() map[string]any {
	paramVals := map[string]any{}

	for name, param := range ifc.Params {
		paramVals[name] = param.GetValue()
	}

	return paramVals
}
