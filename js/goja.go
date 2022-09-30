package js

import (
	"bytes"
	"fmt"
	"github.com/dop251/goja"
	"sync"
	"time"
)

type Script struct {
	pro *goja.Program
}

func NewScript(script string) []byte {
	return []byte(fmt.Sprintf("(function() { %s })();", script))
}

func New(source ...[]byte) (Script, error) {
	pro, err := goja.Compile("script", string(bytes.Join(source, []byte("\n"))), false)
	return Script{pro: pro}, err
}

type Engine struct {
	pool *sync.Pool
}

func NewEngine() *Engine {
	return &Engine{&sync.Pool{
		New: func() interface{} {
			vm := goja.New()
			return vm
		},
	}}
}

func (m *Engine) Execute(s Script, arg interface{}, opts ...ExecOption) (interface{}, error) {
	vm := m.pool.Get().(*goja.Runtime)

	config := &configOptions{
		arg:           arg,
		scriptTimeout: 2 * time.Second,
	}
	for _, o := range opts {
		o(config)
	}
	config.timer = time.AfterFunc(config.scriptTimeout, func() {
		vm.Interrupt("execution timeout")
	})
	defer func() {
		config.timer.Stop()
		vm.ClearInterrupt()
		m.pool.Put(vm)
	}()

	config.set(vm)
	defer config.unset(vm)

	res, err := vm.RunProgram(s.pro)
	if err != nil {
		return nil, castErr(err)
	}
	return res.Export(), nil
}

func castErr(err error) error {
	if exception, ok := err.(*goja.Exception); ok {
		val := exception.Value().Export()
		if castedErr, ok := val.(error); ok {
			return castedErr
		}
	}
	return err
}
