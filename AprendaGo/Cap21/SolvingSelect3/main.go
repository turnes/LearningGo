package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit, 100)

	receive(even, odd, quit)
}

func send(e, o chan int, quit chan bool, x int) {
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	quit <- true
}

func receive(e, o chan int, quit chan bool) {
	for {
		select {
		case v, ok := <-e:
			if ok {
				fmt.Println("even number:", v)
			}
		case v, ok := <-o:
			if ok {
				fmt.Println("odd number:", v)
			}
		case <-quit:
			return
		}
	}
}
