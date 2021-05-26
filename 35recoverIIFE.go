package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured", r)
		} else {
			fmt.Println("Application running perfectly")
		}
	}()

	panic("some error happen")
}
