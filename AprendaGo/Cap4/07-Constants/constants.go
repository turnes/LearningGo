package main

import (
	"fmt"
)

const x = 10

var a int
var b float32

func main() {
	a = x
	b = x

	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)

}
