package js

import (
	"fmt"
	"github.com/dop251/goja"
)

type Runner struct {
	pro *goja.Program
	vm  *goja.Runtime
}

type CallBackFunction func(args ...interface{}) (interface{}, error)

func New(script string) (*Runner, error) {
	pro, err := goja.Compile("", script, true)
	if err != nil {
		return nil, err
	}
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	j := &Runner{pro, vm}
	return j, nil
}

func (j *Runner) Run() (goja.Value, error) {
	return j.vm.RunProgram(j.pro)
}

func (j *Runner) GetObject(name string) (*Object, error) {
	obj := j.vm.Get(name)
	if obj == nil {
		return nil, fmt.Errorf("Got nil value for %s ", name)
	}
	r := obj.ToObject(j.vm)

	return &Object{j, r}, nil
}
func (j *Runner) GetGlobalObject() *Object {
	r := j.vm.GlobalObject()
	return &Object{j, r}
}

func (j *Runner) VM() *goja.Runtime {
	return j.vm
}

func (j *Runner) InjectFn(name string, fn interface{}) {
	j.vm.Set(name, fn)
}
