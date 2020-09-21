package main

import (
	"fmt"
)

type pug int

var x pug
var y int

func main() {

	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v, %T\n", y, y)

}
