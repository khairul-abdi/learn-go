// link https://www.digitalocean.com/community/tutorials/how-to-convert-data-types-in-go#:~:text=You%20can%20convert%20numbers%20to,converted%20into%20a%20string%20value.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.Itoa(12)
	var b int = 2
	fmt.Printf("%q\n", a)
	fmt.Println("hello", b)
}