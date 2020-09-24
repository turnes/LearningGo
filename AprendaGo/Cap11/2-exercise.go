package main

import "fmt"

type person struct {
	name            string
	lastName        string
	icecreamFlavors []string
}

var people = map[string]person{
	"Junior": person{"Neymar", "Junior", []string{"Grape", "Apple"}},
	"Cavani": person{"Edison", "Cavani", []string{"Nutella", "Oreo"}},
}

func main() {

	fmt.Println(people)

	p1 := person{"Lionel", "Messi", []string{"Chocolate", "Strarberry"}}
	p2 := person{"Cristiano", "", []string{"Chocolate", "Strarberry"}}

	fmt.Println(p1)
	for _, flavor := range p1.icecreamFlavors {
		fmt.Printf("\t%v", flavor)
	}

	fmt.Println()

	fmt.Println(p2)
	for _, flavor := range p2.icecreamFlavors {
		fmt.Printf("\t%v", flavor)
	}
	fmt.Println()

	people[p1.lastName] = p1
	people[p2.lastName] = p2

	for _, p := range people {
		fmt.Println(p.name, p.lastName)
		for _, f := range p.icecreamFlavors {
			fmt.Printf("\t%v", f)
		}
		fmt.Println()
	}

}
