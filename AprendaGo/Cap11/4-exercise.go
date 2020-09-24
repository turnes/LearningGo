package main

import "fmt"

func main() {

	a := struct {
		color string
		m     map[int]string
		s     []string
	}{
		color: "yellow",
		m: map[int]string{
			1: "red",
			2: "blue",
			3: "black",
		},
		s: []string{"red, blue, green"},
	}

	fmt.Println(a)

}
