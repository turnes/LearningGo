package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for _, index := range array {
		fmt.Println(index)
	}

	fmt.Printf("Type: %T\n", array)
}
