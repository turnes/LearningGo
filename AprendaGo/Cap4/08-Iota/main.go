package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	_ = iota
	x = iota
	_ = iota
	z = iota
)

const (
	c = iota + 1
	d
	e
	f
)

func main() {
	fmt.Println(a, b, x, z)
	fmt.Println(c, d, e, f)

}
