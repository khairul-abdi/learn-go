package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = ""
	var z int = 4
	fmt.Println(z)
	z += 5
	fmt.Println(z)
	fmt.Println(reflect.TypeOf(a))

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

}
