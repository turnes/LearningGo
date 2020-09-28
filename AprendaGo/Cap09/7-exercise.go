package main

import "fmt"

func main() {

	ss := [][]string{
		[]string{"Roger", "Federer", "Tennis"},
		[]string{"Cristiano", "Ronaldo", "Football"},
		[]string{"Usain", "Bolt", "Sprinter"},
	}

	for _, line := range ss {
		for _, value := range line {
			fmt.Printf("%v\t", value)
		}
		fmt.Println()
	}

}
