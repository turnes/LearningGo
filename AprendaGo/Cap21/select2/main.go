package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	quit := make(chan int)

	go receivedQuit(c, quit)
	send(c, quit)

}

func receivedQuit(c chan int, quit chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println("Received:", <-c)
	}
	quit <- 0
}

func send(c chan int, quit chan int) {
	any := 50
	for {
		select {
		case c <- any:
			any++
		case <-quit:
			return
		}
	}
}
