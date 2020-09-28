package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	// spaces between arguments
	fmt.Println("Hello World", "Teste", 100)
	// fmt.Prinln returns 2 variables.
	numberofbytes, errors := fmt.Println("Hello World", "Teste", 100)
	fmt.Println(numberofbytes, errors)
	// use "_" if want to skip a variable. Eg: number of bytes
	_, errors2 := fmt.Println("Hello World", "Teste", 100)
	fmt.Println(errors2)
}
