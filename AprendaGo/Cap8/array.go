package main

import (
	"fmt"
)

var x [5]int
var y = [6]int{1, 2, 3, 4, 5, 6}

func main() {
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	for i := range x {
		fmt.Printf("Index: %v\t value: %v\n", i, x[i])
	}

	fmt.Printf("\nArray X: %T\nArray Y: %T\n", x, y)

	fmt.Printf("Size of X: %v\tSize of y: %v\n", len(x), len(y))

}
