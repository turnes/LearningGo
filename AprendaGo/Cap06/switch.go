package main

import (
	"fmt"
)

var kind interface{}

func main() {

	x := 5

	switch { // bool
	case x < 5:
		fmt.Println("X é menor que 5")
	case x == 5:
		fmt.Println("X é igual a 5")
	case x > 5:
		fmt.Println("X é maior que 5")
	}

	fmt.Println("\n")

	number := 2

	switch number {
	case 1, 2:
		fmt.Println("Number 1 or 2")
		fallthrough
	case 212121:
		fmt.Println("Result of fallthrough")
	case 3, 4:
		fmt.Println("Number 3 or 4")
	case 5, 6:
		fmt.Println("Number 5 or 6")
	default:
		fmt.Println("Default case")
	}

	fmt.Println("\n")

	kind = "1.0909209209029"

	switch kind.(type) {
	case int:
		fmt.Println("It's an int.")
	case bool:
		fmt.Println("It's a bool.")
	case float64:
		fmt.Println("It's a float.")
	case string:
		fmt.Println("It's a string.")
	}

}
