package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"Bravo", "Echo", "Tango", "Kilo", "Alpha", "Hotel", "Delta"}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)

	si := []int{123, 10, 44, 51, 99}
	fmt.Println(si)
	sort.Ints(si)
	fmt.Println(si)
}
