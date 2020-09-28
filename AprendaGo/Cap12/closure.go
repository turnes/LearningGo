package main

import (
	"fmt"
)

func main() {

	a := i()
	a()
	a()
	a()
	fmt.Println(a())

	b := i()
	b()

	fmt.Println(b())

}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
