package js

import (
	"fmt"
	"github.com/dop251/goja"
	"reflect"
)

type Object struct {
	j   *Runner
	obj *goja.Object
}

func (jo *Object) Call(method string, args ...interface{}) (goja.Value, error) {

	met := jo.obj.Get(method)
	if met == nil {
		return nil, fmt.Errorf("Got nil value for %s ", method)
	}
	var fn goja.Callable
	err := jo.j.vm.ExportTo(met, &fn)
	if err != nil {
		return nil, err
	}
	var vars []goja.Value
	for _, a := range args {
		vars = append(vars, jo.j.vm.ToValue(a))
	}
	return fn(jo.obj, vars...)
}

func (jo *Object) CallReturningObj(method string, args ...interface{}) (*Object, error) {
	v, err := jo.Call(method, args...)
	if err != nil {
		return nil, err
	}
	r := v.ToObject(jo.j.vm)

	return &Object{jo.j, r}, nil
}

func (jo *Object) CallReturningStr(method string, args ...interface{}) (string, error) {
	v, err := jo.Call(method, args...)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func (jo *Object) GetNumber(name string) (int64, error) {
	v := jo.obj.Get(name)
	if v == nil {
		return 0, fmt.Errorf("Got nil value for %s ", name)
	}
	if v.ExportType() != reflect.TypeOf(int64(0)) {
		return 0, fmt.Errorf("The variable %s is not number type", name)
	}
	return v.ToInteger(), nil
}

func (jo *Object) GetString(name string) (string, error) {
	v := jo.obj.Get(name)
	if v == nil {
		return "", fmt.Errorf("Got nil value for %s ", name)
	}

	return v.String(), nil
}

func (jo *Object) GetObject(name string) (*Object, error) {

	obj := jo.obj.Get(name)
	if obj == nil {
		return nil, fmt.Errorf("Got nil value for %s ", name)
	}
	r := obj.ToObject(jo.j.vm)

	return &Object{jo.j, r}, nil
}
