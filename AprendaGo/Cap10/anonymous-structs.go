package main

import "fmt"

func main() {

	x := struct {
		name string
		age  int
	}{
		name: "Cristiano Ronaldo",
		age:  35,
	}

	fmt.Println(x)

}
