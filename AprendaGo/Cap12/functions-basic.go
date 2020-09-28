package main

import (
	"fmt"
)

//func (receiver) identifier(parameters) (returns) { code }

func say(word string) {
	fmt.Println(word)
}

func sum2(x, y int) int {
	return x + y
}

// variadic

func sumMany(x ...int) (int, int) {
	var total int
	for _, n := range x {
		total += n
	}

	return total, len(x)
}

func main() {

	// simple function, no return and no parameters
	say("Hello Functions!")

	// simple function, return and paramenters
	a := 1
	b := 7
	result := sum2(a, b)
	fmt.Println(result)

	// multiple returns
	result, items := sumMany(1, 2, 3, 4, 5)
	fmt.Println(result, items)

	// unfurling a slice
	numbers := []int{10, 20, 30, 40, 50}
	result, items = sumMany(numbers...)
	fmt.Println(result, items)

}
