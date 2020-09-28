package main

import (
	"fmt"
)

var array = [3]int{1, 2, 3}
var slice = []int{1, 2, 3}

func main() {

	fmt.Printf("Type of variable array\t%T\n", array)
	fmt.Printf("Type of variable slice\t%T\n", slice)

}
