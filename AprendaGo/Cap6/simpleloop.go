package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
