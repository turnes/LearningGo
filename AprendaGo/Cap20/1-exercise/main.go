package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Routines", runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println("Routines", runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	fmt.Println("Func1")
	wg.Done()
}

func func2() {
	fmt.Println("Func2")
	wg.Done()
}
