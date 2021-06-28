package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	r := recover()
	if r != nil {
		fmt.Println("PESAN r != nil RECOVER ==> ", r, "TYPE=> ", reflect.TypeOf(r))
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("PESAN r == nil RECOVER ==> ", r, "TYPE=> ", reflect.TypeOf(r))
		fmt.Println("Application running perfectly")
	}
}

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
	}
}
