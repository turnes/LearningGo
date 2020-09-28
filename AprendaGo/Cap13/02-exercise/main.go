package main

import (
	"fmt"
)

func main() {

	total := sum([]int{1, 2, 3, 4, 5, 6}...)

	fmt.Println(total)

	total = sumSlice([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(total)
}

func sum(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}

func sumSlice(x []int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}
