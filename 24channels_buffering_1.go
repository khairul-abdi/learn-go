package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "channel1"
	messages <- "channel2"
	messages <- "channel3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
