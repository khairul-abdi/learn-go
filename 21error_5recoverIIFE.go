package main

import "fmt"

func main() {
	data := []string{"superman", "aquaman", "batman"}

	for _, each := range data {
		func() {
			// recover untuk IIFE dalam perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()
			panic("some error happen")
		}()
	}
}
