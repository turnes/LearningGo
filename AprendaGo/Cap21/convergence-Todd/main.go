package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	converge := make(chan int)

	go send(even, odd)
	go receive(even, odd, converge)

	for v := range converge {
		fmt.Println("Number received:", v)
	}

}

func send(e, o chan int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive(e, o, c chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range e {
			c <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v := range o {
			c <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(c)

}
