package main

import "fmt"

type person struct {
	name string
	age  int
}

type professional struct {
	person
	title  string
	salary int
}

func main() {

	// person who does not work
	p1 := person{"Rafael", 25}

	// person who works
	p2 := professional{
		person: person{"Neymar", 28},
		title:  "Football player",
		salary: 36000000,
	}

	p3 := professional{person{"Cristiano", 35}, "Football player", 30000000}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3.name, p3.age, p3.person.name, p3.title)

}
