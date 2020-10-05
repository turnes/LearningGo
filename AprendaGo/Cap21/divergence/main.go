package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 20)

	go diverge(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}

}

func send(ch chan int, x int) {
	for i := 1; i <= x; i++ {
		ch <- i
	}
	close(ch)
}

func diverge(chout, chin chan int) {
	var wg sync.WaitGroup

	for v := range chout {
		wg.Add(1)
		go func(x int) {
			chin <- work(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(chin)
}

func work(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}
