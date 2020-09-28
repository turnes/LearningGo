package main

import (
	"fmt"
)

func main() {
	x := returnFunction()
	y := x(10)
	fmt.Println(y)

	fmt.Println(returnFunction()(5))
}

func returnFunction() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
