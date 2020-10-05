package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := work("peach")
	ch2 := work("apple")
	channel := converge(ch1, ch2)
	for i := 1; i <= 10; i++ {
		fmt.Println(<-channel)
	}
}

func work(s string) chan string {
	channel := make(chan string)
	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Function: %v\tValue: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}

	}(s, channel)

	return channel
}

func converge(ch1, ch2 chan string) chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- <-ch1
		}
	}()

	go func() {

		for {
			channel <- <-ch2
		}
	}()

	return channel

}
