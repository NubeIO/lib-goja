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

function sum(a) {
  return a
}

function sum2(a) {
  return a
}


var aa = sum()
var bb = sum2()
	

	`
	j, err := js.New(code)

	if err != nil {
		fmt.Printf("Error loading JS code %v", err)
		return
	}
	//j.InjectFn("add", add)
	//j.InjectFn("sub", sub)
	//j.GetGlobalObject()

	_, err = j.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	//sobj, err := j.GetObject("sum")

	call, err := j.CallFunction("sum", 332)
	fmt.Println(err)
	fmt.Println(call)
	if err != nil {
		return
	}
	call2, err := j.CallFunction("sum2", 4444)
	fmt.Println(err)
	fmt.Println(call2)
	if err != nil {
		return
	}

	fmt.Println()
	//res2, err := sobj.Call("number", 44)
	//fmt.Println(err)
	//restr := res2.String()
	//fmt.Println(restr)
	//if err != nil {
	//	return
	//}

	res, err := j.GetGlobalObject().GetNumber("aa")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Printf("value returned from JS %d", res)

	res, err = j.GetGlobalObject().GetNumber("bb")
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
