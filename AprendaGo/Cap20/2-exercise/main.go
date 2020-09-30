package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hi! I'm", p.first, p.last)
}

type human interface {
	speak()
}

func speakSomething(h human) {
	h.speak()
}

func main() {

	p1 := person{"Cristiano", "Ronaldo", 35}

	speakSomething(&p1)

}
