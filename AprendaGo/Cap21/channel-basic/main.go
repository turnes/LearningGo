package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 1) // biderecional
	channel <- 42
	fmt.Println(<-channel)

	fmt.Println("-----")

	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send
	fmt.Printf("Channel Bidirecional:\t%T\n", channel)
	fmt.Printf("Channel only receives:\t%T\n", cr)
	fmt.Printf("Channel only sends:\t%T\n", cs)

	// change to general to specific
	fmt.Println("-----")
	fmt.Printf("Channel Bidirecional to receive one:\t%T\n", (<-chan int)(channel))
	fmt.Printf("Channel Bidirecional to send one:\t%T\n", (chan<- int)(channel))

}
