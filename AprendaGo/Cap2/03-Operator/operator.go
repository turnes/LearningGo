package main

import (
	"fmt"
)

var y = "Teste" // package level scope

func main() {
	x := 10 // gopher or punisher

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x: %v, %T\n", z, z)

	result := 10 > 20
	fmt.Printf("result: %v, %T\n", result, result)

}
