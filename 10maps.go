package main

import (
	"fmt"
	"reflect"
)

func main()  {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println(reflect.TypeOf(m))

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("delete k2 result map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("maps n: ", n)
}