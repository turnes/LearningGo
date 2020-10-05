package main

import "fmt"

func main() {
	ch := make(chan int) // create channel
	times := 100         // how many numbers will be sent to the channel

	go send(ch, times) // go routine to send the numbers to the channel

	receive(ch) // remove the number from the channel

}

func send(ch chan int, times int) { // function that sends numbers to the channel
	for i := 1; i <= times; i++ {
		ch <- i // sending number "i" to the channel
	}

	close(ch) // indicates that no more values will be sent. It is utils when you use range.
}

func receive(ch chan int) { // function that removes numbers from the channel
	for v := range ch { // remove an element from the channel
		fmt.Println(v) // print the value of the element.
	}

}
