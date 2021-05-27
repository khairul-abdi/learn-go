package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Input data: ")
	fmt.Scan(&input)

	num, err := strconv.Atoi(input)

	if err != nil {
		// handle error
		fmt.Println(input, "isn`t number")
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(num, " is number")
}
