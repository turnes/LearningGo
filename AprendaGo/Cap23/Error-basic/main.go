package main

import (
	"fmt"
)

func main() {

	result, err := fmt.Println("Hello")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}
