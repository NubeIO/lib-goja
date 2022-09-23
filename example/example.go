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
var out1 =  add(33, 33)
out1 = sub(out1, 10)
	`
	j, err := js.New(code)

	if err != nil {
		fmt.Printf("Error loading JS code %v", err)
		return
	}
	j.InjectFn("add", add)
	j.InjectFn("sub", sub)
	j.GetGlobalObject()

	_, err = j.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := j.GetGlobalObject().GetNumber("out1")
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
