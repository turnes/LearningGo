package main

import (
	"fmt"
)

func main() {
	x := `"Literal string \n\n\n does not work
\"True\"`
	fmt.Println(x)
}
