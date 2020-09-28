package main

import (
	"fmt"
)

//func (receiver) identifier(parameters) (returns) { code }

type person struct {
	name     string
	lastName string
	age      int
}

func (p person) toString() {
	fmt.Printf("Name: %v\tLast Name: %v\tAge: %v\n", p.name, p.lastName, p.age)
}

func main() {

	p1 := person{"Cristiano", "Ronaldo", 35}
	p1.toString()

}
