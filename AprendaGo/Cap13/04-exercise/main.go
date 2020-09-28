package main

import (
	"fmt"
)

type person struct {
	name  string
	lname string
	age   int
}

func (p person) toString() {
	fmt.Printf("Name: %v %v\tAge: %v\n", p.name, p.lname, p.age)
}

func main() {
	p1 := person{"Neymar", "Junior", 28}
	p1.toString()
}
