package main

import (
	"fmt"
)

func main() {
	s := "Hello, playground"
	fmt.Printf("%v\t%T", s, s)

	for _, v := range s {
		fmt.Printf("\n%v\t%T\t%#U\t%#x", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("\n%v\t%T\t%#U\t%#x", s[i], s[i], s[i], s[i])
	}

	sb := []byte(s)
	fmt.Printf("\n%v\t%T", sb, sb)

	for _, v := range sb {
		fmt.Printf("\n%v\t%T\t%#U\t%#x", v, v, v, v)
	}

}
