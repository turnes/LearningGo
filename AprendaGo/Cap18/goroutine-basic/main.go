package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GoRoutine:", runtime.NumGoroutine())

	wg.Add(2)

	go routine1()
	go routine2()

	fmt.Println("GoRoutine:", runtime.NumGoroutine())
	wg.Wait()

}

func routine1() {
	for i := 0; i < 100; i++ {
		fmt.Println("routine1:", i)
		//time.Sleep(20)
	}
	wg.Done()
}

func routine2() {
	for i := 100; i > 0; i-- {
		fmt.Println("routine2:", i)
		//time.Sleep(20)
	}
	wg.Done()
}
