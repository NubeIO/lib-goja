package js

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
)

func TestGoja_AddFunction(t *testing.T) {
	a := assert.New(t)

	const SCRIPT = `
	var response = {};
	try {
		response["test"] = f("string", "unknown", "test");
	} catch (e) {
		if (!(e instanceof GoError)) {
			throw(e);
		}
		if (e.value.Error() !== "TEST") {
			throw("Unexpected value: " + e.value.Error());
		}
	}
	return response;
	`

	f := func(varchar string, integer int, object string) (interface{}, error) {
		a.Equal("string", varchar)
		a.Equal(0, integer)
		a.Equal("test", object)
		return "test", nil
	}
	vm := goja.New()
	vm.Set("f", f)
	resp, err := vm.RunString(fmt.Sprintf("(function() { %s })();", SCRIPT))
	a.NoError(err)
	a.Equal(resp.Export(), map[string]interface{}{"test": "test"})

	f2 := func(varchar string, integer int, object string) (interface{}, error) {
		a.Equal("string", varchar)
		a.Equal(0, integer)
		a.Equal("test", object)
		return "test", errors.New("TEST")
	}
	vm = goja.New()
	vm.Set("f", f2)
	resp, err = vm.RunString(fmt.Sprintf("(function() { %s })();", SCRIPT))
	a.NoError(err)
	a.Equal(resp.Export(), map[string]interface{}{})
}

func TestScript_Default(t *testing.T) {
	a := assert.New(t)

	const SHARED = `
	function shared() {
		return arg["key"] + 2
	}
`
	const SCRIPT = `
	return shared()
`
	script, err := New([]byte(SHARED), []byte(fmt.Sprintf("(function() { %s })();", SCRIPT)))
	a.NoError(err)

	arg := map[string]interface{}{"key": 3, "4": 7}

	result, err := NewEngine().Execute(script, arg)
	a.NoError(err)

	a.Equal(int64(5), result)
}

func TestScript_WithLogging2(t *testing.T) {
	a := assert.New(t)

	const SCRIPT = `
	console.log(arg)
	console.log(1, 2)
	console.log("test")
	return arg
`
	script, err := New([]byte(fmt.Sprintf("(function() { %s })();", SCRIPT)))
	a.NoError(err)

	arg := map[string]interface{}{"key": 3}

	logBuf := new(bytes.Buffer)
	result, err := NewEngine().Execute(script, arg, WithLogging(logBuf))
	fmt.Println(logBuf.String())
	fmt.Println(result)

	//a.NoError(err)
	//a.Equal("[{\"key\":3}],\n[1,2,3],\n[\"test\"],\n", logBuf.String())
	//a.Equal(int64(5), result)
}

func TestScript_WithLogging(t *testing.T) {
	a := assert.New(t)

	const SCRIPT = `
	console.log(arg)
	console.log(1, 2, 3)
	console.log("test")
	return 5
`
	script, err := New([]byte(fmt.Sprintf("(function() { %s })();", SCRIPT)))
	a.NoError(err)

	arg := map[string]interface{}{"key": 3}

	logBuf := new(bytes.Buffer)
	result, err := NewEngine().Execute(script, arg, WithLogging(logBuf))
	a.NoError(err)
	a.Equal("[{\"key\":3}],\n[1,2,3],\n[\"test\"],\n", logBuf.String())
	a.Equal(int64(5), result)
}

func GetInt(in interface{}) (val int) {
	switch i := in.(type) {
	case int:
		val = i
	case float64:
		val = int(i)
	case float32:
		val = int(i)
	case int64:
		val = int(i)
	default:
		return 0
	}
	return val
}

func TestScript_WithData2(t *testing.T) {
	a := assert.New(t)

	const SCRIPT = `
	function func() {
		return arg["arg1"] * arg["arg2"]
	}
	
	var a = in1 + in2
	var getOutValue = {
"priority": {
"_1": null,
"_2": null,
"_3": null,
"_4": null,
"_5": null,
"_6": null,
"_7": null,
"_8": null,
"_9": null,
"_10": null,
"_11": null,
"_12": null,
"_13": func(),
"_14": a,
"_15": in1,
"_16": in2
}
}
let out = [getOutValue.priority["_15"], 33]
	return out
`
	script, err := New([]byte(fmt.Sprintf("(function() { %s })();", SCRIPT)))
	a.NoError(err)

	arg := map[string]interface{}{"arg1": 2, "arg2": 2}

	result, err := NewEngine().Execute(script, arg,
		WithSet("in1", 22),
		WithSet("in2", 22))

	fmt.Println(result)

}

func TestScript_WithData(t *testing.T) {
	a := assert.New(t)

	const SCRIPT = `
	return i + str + mp[3]+arg["key"]+arr
`
	script, err := New([]byte(fmt.Sprintf("(function() { %s })();", SCRIPT)))
	a.NoError(err)

	arg := map[string]interface{}{"key": 3}

	result, err := NewEngine().Execute(script, arg,
		WithSet("i", 1),
		WithSet("str", "two"),
		WithSet("mp", map[string]interface{}{"3": "four"}),
		WithSet("arr", []int{5, 6, 7}))
	fmt.Println(result)
	a.NoError(err)

	a.Equal("1twofour35,6,7", result)
}

func TestScript_WithFunc(t *testing.T) {
	a := assert.New(t)

	const SCRIPT = `
	return sqrt(arg["key"])
`
	script, err := New([]byte(fmt.Sprintf("(function() { %s })();", SCRIPT)))
	a.NoError(err)

	arg := map[string]interface{}{"key": 3}

	sqrt := func(x int) int {
		return x * x
	}
	result, err := NewEngine().Execute(script, arg, WithSet("sqrt", sqrt))
	a.NoError(err)

	a.Equal(int64(9), result)
}

func TestScript_WithDataWithFunc(t *testing.T) {
	a := assert.New(t)

	const SCRIPT = `
	return sqrt(arg["key"]) + sqrt(i)
`
	script, err := New([]byte(fmt.Sprintf("(function() { %s })();", SCRIPT)))
	a.NoError(err)

	arg := map[string]interface{}{"key": 3}

	sqrt := func(x int) int {
		return x * x
	}
	result, err := NewEngine().Execute(script, arg, WithSet("sqrt", sqrt), WithSet("i", 1))
	a.NoError(err)

	a.Equal(int64(10), result)
}
