package main

import (
	"fmt"
	"sort"
)

type car struct {
	name         string
	power        float64 // PS
	fuel         string  // petrol , diesel, eletric
	consuption   float64 // Km/l
	acceleration float64 //0 - 100 Km/h

}

func printCars(cars []car, msg string) {
	fmt.Println("\n", msg)
	for _, c := range cars {
		fmt.Println(c)
	}
}

type sortByConsuption []car

func (c sortByConsuption) Len() int           { return len(c) }
func (c sortByConsuption) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c sortByConsuption) Less(i, j int) bool { return c[i].consuption > c[j].consuption }

type sortByAcceleration []car

func (c sortByAcceleration) Len() int           { return len(c) }
func (c sortByAcceleration) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c sortByAcceleration) Less(i, j int) bool { return c[i].acceleration > c[j].acceleration }

type sortByPower []car

func (c sortByPower) Len() int           { return len(c) }
func (c sortByPower) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c sortByPower) Less(i, j int) bool { return c[i].power < c[j].power }

func main() {

	cars := []car{
		car{"a", 90.1, "petrol", 15.4, 9.1},
		car{"b", 60.5, "diesel", 18.9, 15.3},
		car{"c", 120.3, "eletric", 20, 6.4},
		car{"d", 100, "petrol", 13.5, 8.7},
	}

	printCars(cars, "Original")
	sort.Sort(sortByConsuption(cars))
	printCars(cars, "Sorted by Consuption")

	sort.Sort(sortByAcceleration(cars))
	printCars(cars, "Sorted by Acceleration")

	sort.Sort(sortByPower(cars))
	printCars(cars, "Sorted by Power")

}
