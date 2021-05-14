package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("go" + "lang")

	var a float64
	a = 1
	var b = 0.2
	var c, d int = 1, 2

	f := 2

	fmt.Println(a + b)
	fmt.Println(c + d)
	fmt.Println(f)

	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// fmt.Println( reflect.TypeOf("3" + 4))
	fmt.Println("3" + 4)

}
