package main

import "fmt"

func main() {

	pizzaFlavors := []string{"portuguesa", "calabresa", "frango", "pepperoni", "quatro queijos"}
	fmt.Println("\n", pizzaFlavors, "\n")

	mySlices := pizzaFlavors[:2]

	fmt.Println(mySlices, "\n")

	for flavor := 0; flavor < len(pizzaFlavors); flavor++ {
		fmt.Println(pizzaFlavors[flavor])
	}

	//remove frango
	pizzaFlavors = append(pizzaFlavors[:2], pizzaFlavors[3:]...)

	fmt.Println("\n", pizzaFlavors, "\n")

	for _, flavor := range pizzaFlavors {
		fmt.Println(flavor)
	}
}
