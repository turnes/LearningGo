package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Routines", runtime.NumGoroutine())
	counter := 0
	goroutines := 1000

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Routines", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("Routines", runtime.NumGoroutine())

}
