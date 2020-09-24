package main

import (
	"fmt"
)

func main() {

	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(slice)
	slice = append(slice, 52)
	fmt.Println(slice)
	slice = append(slice, 53, 54, 55)
	fmt.Println(slice)

	slice2 := []int{56, 57, 58, 59, 60}

	slice = append(slice, slice2...)
	fmt.Println(slice)

}
