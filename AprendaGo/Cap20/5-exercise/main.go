package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

const numberOfGoroutines = 100

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())

	counter = 0

	for i := 1; i <= numberOfGoroutines; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()

		fmt.Println("Goroutine", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}
