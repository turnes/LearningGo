package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c, 10)

	for v := range c {
		fmt.Println("Received from channel:", v)
	}

}

func send(s chan<- int, t int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}
