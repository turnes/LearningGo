package main

import (
	"fmt"
)

func main() {

	slice := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	fmt.Printf("\tThe first three items\n")

	for _, value := range slice[:3] {
		fmt.Println(value)
	}

	fmt.Printf("\tFrom item 5 to the end\n")

	for _, value := range slice[4:] {
		fmt.Println(value)
	}

	fmt.Printf("\tFrom item 2 to item 7\n")

	for _, value := range slice[1:7] {
		fmt.Println(value)
	}

	fmt.Printf("\tFrom item 2 to item 7\n")

	for _, value := range slice[2 : len(slice)-1] {
		fmt.Println(value)
	}

}
