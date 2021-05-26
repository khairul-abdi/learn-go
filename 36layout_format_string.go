//link : https://dasarpemrogramangolang.novalagung.com/A-string-format.html
package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	fmt.Printf("%b\n", data.age) //Angka jadi biner:  11010
}
