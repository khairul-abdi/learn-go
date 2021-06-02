package main

import (
	"fmt"

	"A.26.7-golang-level-akses-init/library"
)

func main() {
	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
