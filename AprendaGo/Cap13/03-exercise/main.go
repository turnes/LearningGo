package main

import (
	"fmt"
)

func main() {

	x := 10
	defer fmt.Println("Defer: ", x+1)
	x += 5
	fmt.Println(x)

}
