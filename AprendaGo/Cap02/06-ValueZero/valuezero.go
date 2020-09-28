package main

import (
	"fmt"
)

var a int
var b float32
var c string
var d bool

func main() {

	// default values of each type
	fmt.Printf("type: %T, value: %v\n", a, a)
	fmt.Printf("type: %T, value: %v\n", b, b)
	fmt.Printf("type: %T, value: %v\n", c, c)
	fmt.Printf("type: %T, value: %v\n", d, d)

}
