package main
import "fmt"
func main()  {
	
	var message1 = "Nama saya 'Khairul Abdi'. \n Salam kenal.\n Mari belajar 'Golang'."

var message2 = `Nama saya "Khairul Abdi".
Salam kenal.
Mari belajar "Golang".`

	fmt.Println(message1)
	fmt.Println(message2)
}