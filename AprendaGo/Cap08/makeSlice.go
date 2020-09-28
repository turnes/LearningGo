package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)
	fmt.Printf("len: %v\tcap: %v\t\n", len(slice), cap(slice))

	slice = append(slice, 6, 7, 8, 9, 10, 11)
	fmt.Println(slice)
	fmt.Printf("len: %v\tcap: %v\t\n", len(slice), cap(slice))

}
