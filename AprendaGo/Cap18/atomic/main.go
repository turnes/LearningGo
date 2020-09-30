package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Routines", runtime.NumGoroutine())
	var counter int64
	counter = 0
	goroutines := 1000

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//runtime.Gosched()
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Routines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
	fmt.Println("Routines", runtime.NumGoroutine())

}
