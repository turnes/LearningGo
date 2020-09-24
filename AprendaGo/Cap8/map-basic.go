package main

import "fmt"

func main() {

	amigos := map[string]int{
		"alfredo": 121212,
		"joana":   24034093,
		"rafael":  322323,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])
	amigos["gopher"] = 73823877
	fmt.Println(amigos["gopher"])

	// comma ok idiom

	fantasma, ok := amigos["fantasma"]
	fmt.Println(fantasma, ok)

	if exist, ok := amigos["ghost"]; ok {
		fmt.Println(exist)
	}

	if exist, ok := amigos["rafael"]; ok {
		fmt.Println(exist)
	}

}
