package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const numberOfGoroutines = 100

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())
	counter := 0

	for i := 1; i <= numberOfGoroutines; i++ {
		wg.Add(1)
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutine", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}
