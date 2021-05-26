// Link https://dasarpemrogramangolang.novalagung.com/A-tipe-data.html
package main

import ( "fmt"; "reflect" )

func main()  {
	var positiveNumber uint8 	= 89
	var negativeNumber 				= -1243423644

// fmt.Printf("bilangan positif: %d\n", reflect.TypeOf(positiveNumber))
// fmt.Printf("bilangan negatif: %d\n", reflect.TypeOf(negativeNumber))
fmt.Println("bilangan positif: ", reflect.TypeOf(positiveNumber))
fmt.Println("bilangan negatif: ", reflect.TypeOf(negativeNumber))
}