package main

import "fmt"

type person struct {
	name     string
	lastName string
	address  string
}

func main() {

	p1 := person{"Cristiano", "Ronaldo", "Juventus Stadium"}

	fmt.Println(p1)

	changeAddress(&p1, "the Stadium")
	fmt.Println(p1)

}

func changeAddress(p *person, adress string) {
	(*p).address = adress
	//p.address = address
}
