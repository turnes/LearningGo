package main

import "fmt"

func main() {

	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{6, 7, 8, 9, 10}

	fmt.Println(sliceA)
	fmt.Println(sliceB)

	sliceC := append(sliceA, sliceB...) //... -> unfurl or enumerations

	fmt.Println(sliceC)

	sliceC = append(sliceC, 11, 12, 13, 14, 15)
	fmt.Println(sliceC)

}
