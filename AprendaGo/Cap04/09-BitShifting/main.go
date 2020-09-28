package main

import (
	"fmt"
)

func main() {
	x := 2
	y := x >> 1
	z := x << 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
}
