package main

import (
	"fmt"
)

func main() {

	slice := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("Type: %T\n", slice)

}
