package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Routines", runtime.NumGoroutine())
	counter := 0
	goroutines := 1000

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Routines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
	fmt.Println("Routines", runtime.NumGoroutine())

}
