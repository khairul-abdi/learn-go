package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		//Cara 1
		// return false, errors.New("cannot be empty")

		//Cara 2
		return false, fmt.Errorf("cannot be empty")
	}
	return true, nil
}
