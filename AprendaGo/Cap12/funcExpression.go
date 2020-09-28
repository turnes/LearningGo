package main

import (
	"fmt"
)

func main() {

	x := 5
	y := func(x int) {
		fmt.Println(x, "times 10 is:")
		fmt.Println(x * 10)
	}
	y(x)

}
