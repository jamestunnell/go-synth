package synth

import "reflect"

type Interface struct {
	Inputs   map[string]Input
	Controls map[string]Control
	Params   map[string]Param
	Outputs  map[string]Output
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
	ifc.extract(reflect.ValueOf(b))
}

func (ifc *Interface) extract(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return
	}

	vt := v.Type()

	for i := 0; i < vt.NumField(); i++ {
		sf := vt.Field(i)

		if !sf.IsExported() {
			continue
		}

		fv := v.Field(i)

		if !fv.CanInterface() {
			continue
		}

		switch vv := fv.Interface().(type) {
		case Control:
			ifc.Controls[sf.Name] = vv
		case Input:
			ifc.Inputs[sf.Name] = vv
		case Param:
			ifc.Params[sf.Name] = vv
		case Output:
			ifc.Outputs[sf.Name] = vv
		default:
			// try to extract from a possible embedded struct
			ifc.extract(fv)
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
