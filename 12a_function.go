// link https://www.golangprograms.com/passing-multiple-string-arguments-to-a-function.html
package main

import "fmt"

func main() {
	
	variadicExample()
	variadicExample("red", "blue")
	variadicExample("red", "blue", "green")
	variadicExample("red", "blue", "green", "yellow")
	variadicExample(8)
}

func variadicExample(s ...string) {
	fmt.Println(s)
}