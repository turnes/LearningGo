package main

import (
	"fmt"
)

type pug int

var x pug

func main() {

	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Println(x)

}
