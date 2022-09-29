package main

import (
	"fmt"
	"github.com/NubeIO/lib-goja/js"
)

func Test() {
	code := `
var a = passInValue
var getOutValue = ""
if(a=="hello") {
getOutValue = "hello"
} else {
getOutValue = "not"
}
	`
	j, err := js.New(code)

	if err != nil {
		fmt.Printf("Error loading JS code %v", err)
		return
	}
	j.Set("passInValue", "hell")

	j.GetGlobalObject()

	_, err = j.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := j.GetGlobalObject().GetString("getOutValue")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Printf("value returned from JS %s", res)

}

func main() {
	Test()
}
