package main

import (
	"fmt"
)

const (
	_ = 2000 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
