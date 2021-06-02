package main

import (
	"fmt"

	"A.26.4-golang-level-akses-exported-dan-unexported-struct-dan-property/library"
)

func main() {
	var s1 = library.Student{"ethan", 21}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)
}
