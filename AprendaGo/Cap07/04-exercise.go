package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1985

	for {
		if x > time.Now().Year() {
			break
		}
		fmt.Println(x)
		x++
	}

}
