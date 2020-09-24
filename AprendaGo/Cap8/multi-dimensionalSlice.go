package main

import "fmt"

func main() {

	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 6, 7},
		[]int{7, 8, 9},
		[]int{10, 11, 12, 13},
	}

	fmt.Printf("Type: %T\tValue:\n\t\t%v\n", ss, ss)

	for x := range ss {
		for y := range ss[x] {
			fmt.Println(ss[x][y])
		}
	}
}
