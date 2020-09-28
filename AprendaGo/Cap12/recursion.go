package main

import (
	"fmt"
)

func main() {

	fmt.Println(factorial(4))
	fmt.Println(loops(6))

}

func factorial(x int) int {
	if x == 1 {
		return x
	}
	return x * factorial(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
