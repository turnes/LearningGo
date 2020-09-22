package main

import (
	"fmt"
)

func main() {
	a := 1 == 0
	b := 1 != 0
	c := 1 <= 0
	d := 1 < 0
	e := 1 >= 0
	f := 1 > 0

	fmt.Printf(`
	1 == 0	%v
	1 != 0	%v
	1 <= 0 	%v
	1 < 0 	%v
	1 >= 0 	%v
	1 > 0 	%v `, a, b, c, d, e, f)
}
