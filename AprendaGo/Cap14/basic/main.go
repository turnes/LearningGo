package main

import (
	"fmt"
)

func main() {
	x := 10

	y := &x

	fmt.Printf("X: %T\tY: %T\n", x, y)
	fmt.Printf("X: %v\tY: %v | %v\n", x, y, *y)

	x = 20
	fmt.Printf("X: %v\tY: %v | %v\n", x, y, *y)
	*y++
	fmt.Printf("X: %v\tY: %v | %v\n", x, y, *y)
}
