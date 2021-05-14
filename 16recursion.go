package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		fmt.Println("Masuk")
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(4))
}