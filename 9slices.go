package main

import (
	"fmt"
	"reflect"
)

func main()  {
	s := make([]string, 3)
	q := make([]int, 3)
	fmt.Println("emp: ", s)
	fmt.Println("emp q =>: ", q)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	
	fmt.Println("emp: ", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd s:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy c:", c)

	d := s
	fmt.Println("cpy d:", d)

	d[5] = "z"
	fmt.Println("cpy d:", reflect.TypeOf(d))

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
    fmt.Println("twoD P: ", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		fmt.Println("twoD 1 : ", twoD)
		for j := 0; j < innerLen; j++ {
		    fmt.Println("twoD 2 : ", twoD)
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)



	arr := make([]int, 3)
    fmt.Println("arr P: ", arr)
	for i := 0; i < 3; i++ {
	}
	fmt.Println("arr: ", arr)
	
}