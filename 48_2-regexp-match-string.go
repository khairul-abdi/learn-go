package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "Bbanana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
	// true
}
