package main

import "fmt"

type person struct {
	name            string
	lastName        string
	icecreamFlavors []string
}

func main() {

	p1 := person{"Lionel", "Messi", []string{"Chocolate", "Strarberry"}}
	p2 := person{"Cristiano", "Ronaldo", []string{"Vanilla", "Pistache"}}

	fmt.Println(p1.name, p1.lastName)

	for _, flavor := range p1.icecreamFlavors {
		fmt.Printf("\t%v", flavor)

	}

	fmt.Println("\n", p2.name, p2.lastName)

	for _, flavor := range p2.icecreamFlavors {
		fmt.Printf("\t%v", flavor)
	}
	fmt.Println()

}
