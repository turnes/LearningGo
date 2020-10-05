package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	go send(ch)
	receive(ch)

	// for v := range ch {
	// 	fmt.Printf("Routine %v:\t%v\n", v/100, v)
	// }
}

func send(c chan int) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			for j := 0; j < 10; j++ {
				c <- (x * 100) + j
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(c)
}

func receive(c chan int) {
	for v := range c {
		fmt.Printf("Routine %v:\t%v\n", v/100, v)
	}
}
