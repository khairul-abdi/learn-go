package main

import "fmt"

func main() {

	messages := make(chan int)

	go func() {
		messages <- 23
	}()

	msg := <-messages
	fmt.Println(msg)
}
