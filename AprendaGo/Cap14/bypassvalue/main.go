package main

import (
	"fmt"
)

func main() {

	x := 10
	passValue(x)
	fmt.Println("Value of x after function:", x)
	fmt.Println("---------")
	passPointer(&x)
	fmt.Println("Value of x after function:", x)

}

func passValue(number int) {
	number++
	fmt.Println("Function passValue:", number)

}

func passPointer(pointer *int) {
	*pointer++
	fmt.Println("Function passPointer:", *pointer)

}
