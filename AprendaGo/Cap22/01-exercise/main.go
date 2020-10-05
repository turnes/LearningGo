package main

import (
	"fmt"
)

func main() {

	// anonymous function
	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	fmt.Println("----------")

	//buffer
	c = make(chan int, 1)
	c <- 44
	fmt.Println(<-c)

}
