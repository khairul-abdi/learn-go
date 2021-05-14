// link https://www.golangprograms.com/how-to-convert-string-to-integer-type-in-go.html#:~:text=In%20order%20to%20convert%20string,can%20use%20the%20following%20methods.&text=You%20can%20use%20the%20strconv,the%20error%20(if%20any).
package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}