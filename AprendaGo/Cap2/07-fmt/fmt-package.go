package main

import (
	"fmt"
)

func main() {
	//interpreted string literals
	x := "Hello!\nHow are you ?\tI hope you \"are\" fine."
	fmt.Println(x)

	// raw  string literals
	y := `"Hello!\nHow are you ?\tI hope you \"are\" fine."`
	fmt.Println(y)

	a := "Hello, "
	b := "fmt package!"

	//PRINT
	fmt.Print(a)
	fmt.Print(b)
	fmt.Println(a)
	fmt.Println(b)
	//SPRINT
	c := fmt.Sprint(a, b)
	fmt.Println(c)
	d, e := 2, 4
	f := fmt.Sprint(d, e) // adds space between operands (except strings)
	fmt.Println(f)

	//FPRINT

}
