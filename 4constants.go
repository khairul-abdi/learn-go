package main

import (
	"fmt"
	"math"
)

const s string = "constant"
var t string = "constant"

func main()  {
	fmt.Println(s)
	fmt.Println(t)

	var aa string
	aa = "hello"
	fmt.Println(aa)

	const n = 50000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println("PRINT 1: ",int64(d))

	const f = 3e20 - n
	fmt.Println(f)
	// fmt.Println("PRINT 2: ",int64(f))

	fmt.Println(math.Sin(n))

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Println(negativeNumber)
}