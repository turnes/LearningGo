package main

import (
	"fmt"
)

func main() {
	var i uint16 //65535
	// i = 65536 --> constant 65536 overflows uint16
	i = 65535
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)

}
