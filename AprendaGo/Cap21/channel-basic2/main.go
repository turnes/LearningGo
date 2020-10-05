package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c, 10)

	receive(c)

}

func send(s chan<- int, v int) {
	fmt.Println("Sending:", v)
	s <- v
}

func receive(r <-chan int) {
	fmt.Println("Receiving: ", <-r)

}
