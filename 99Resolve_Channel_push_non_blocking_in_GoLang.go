// link : https://dev.to/calj/channel-push-non-blocking-in-golang-1p8g

package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	for i := 0; i < 10; i++ {
		select {
		case ch <- "hello":
			fmt.Println("Pushed a message to the channel")
		default:
			fmt.Println("WARN: The channel is full")
		}
	}
}
