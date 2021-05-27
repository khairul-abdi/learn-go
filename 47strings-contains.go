package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)
}
