package main

import (
	"bytes"
	"fmt"
	"github.com/NubeIO/lib-goja/js"
)

func main() {

	const eg = `let pri = {
    "priority": {
            "_14": Number(arg["in1"]),
            "_15": Number(arg["in2"]),
            "_16": Number(arg["in3"])
        
        }
}

return JSON.stringify(pri)


`
	script, err := js.New(js.NewScript(eg))
	if err != nil {
		return
	}

	arg := map[string]interface{}{"in1": 22.5, "in2": 33.3, "in3": nil}

	consoleLogs := new(bytes.Buffer)
	result, err := js.NewEngine().Execute(script, arg, js.WithLogging(consoleLogs))

	fmt.Println(consoleLogs)
	fmt.Println(result)
}
