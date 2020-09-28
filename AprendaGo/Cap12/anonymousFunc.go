package main

import (
	"fmt"
)

// related to go routines

func main() {
	x := 5

	func(x int) {
		fmt.Println(x, "times 10 is :")
		fmt.Println(x * 10)
	}(x)

}
