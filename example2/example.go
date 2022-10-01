package main

import (
	"bytes"
	"fmt"
	"github.com/NubeIO/lib-goja/js"
)

func main() {

	const SCRIPT = `
	console.log(arg)
	console.log(1, 2, arg["key"])
	var x = Number(arg["key"])
	console.log("test")
	return [1,2,x]
`
	script, err := js.New(js.NewScript(SCRIPT))
	if err != nil {
		return
	}

	arg := map[string]interface{}{"key": "22"}

	consoleLogs := new(bytes.Buffer)
	result, err := js.NewEngine().Execute(script, arg, js.WithLogging(consoleLogs))

	fmt.Println(consoleLogs)
	fmt.Println(result)
	arr, ok := result.([]interface{})
	fmt.Println(arr, ok)
	fmt.Println(arr[2].(int64) * 2)
}
