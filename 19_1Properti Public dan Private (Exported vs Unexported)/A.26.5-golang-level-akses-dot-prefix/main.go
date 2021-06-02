package main

import (
	"fmt"

	. "A.26.5-golang-level-akses-dot-prefix/library"
)

func main() {
	var s1 = Student{"ethan", 21}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)
}
