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

	go send(ch1, 100)

	go diverge(ch1, ch2, 5)

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

func diverge(chout, chin chan int, times int) {
	var wg sync.WaitGroup

	for i := 1; i <= times; i++ {
		wg.Add(1)
		go func() {
			for v := range chout {
				chin <- work(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(chin)
}

func work(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}
