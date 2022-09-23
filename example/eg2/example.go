package main

import (
	"fmt"
	"github.com/NubeIO/lib-goja/js"
)

func add(a, b int) (int, error) {
	return a + b, nil
}

func sub(a, b int) (int, error) {
	return a - b, nil
}

func Test() {
	code := `
	var getOutValue = passInValue*2
	`
	j, err := js.New(code)

	if err != nil {
		fmt.Printf("Error loading JS code %v", err)
		return
	}
	j.Set("passInValue", 22.465645)

	j.GetGlobalObject()

	_, err = j.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := j.GetGlobalObject().ToInteger("getOutValue")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Printf("value returned from JS %d", res)

}

func main() {
	Test()
}
