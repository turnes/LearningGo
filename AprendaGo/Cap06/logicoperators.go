package main

import (
	"fmt"
)

func main() {
	x := 2
	y := 4
	z := 6
	if x == 2 || x == 3 {
		fmt.Println("X is equal 2 or 3")
	}

	if x == 2 && y == 4 {
		fmt.Println("X is equal 2 or 4")
	}

	if z%2 == 0 && z%3 == 0 {
		fmt.Println("Z is divisible by 6")
	}

}
