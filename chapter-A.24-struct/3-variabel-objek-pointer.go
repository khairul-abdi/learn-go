package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{name: "wick", grade: 2}
	var s11 = student{name: "diah", grade: 3}

	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
	fmt.Println("student 4, name :", s11.name)
	fmt.Println("============================")

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
	fmt.Println("student 4, name :", s11.name)

	s2.name = "diah"
	fmt.Println("============================")
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
	fmt.Println("student 4, name :", s11.name)

}
