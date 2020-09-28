package main

import (
	"fmt"
	"reflect"
)

//interface
type person struct {
	name     string
	lastName string
	age      int
}

type dentist struct {
	person
	salary      int
	specialties []string
}

type architect struct {
	person
	kind     string
	projects []string
}

func (d dentist) introduceYourself() {
	fmt.Printf("Hi! My name is %v %v and I'm %v years old. My profession is %v. My specialties are: %v\n", d.name, d.lastName, d.age, reflect.TypeOf(d).Name(), d.specialties)

}

func (a architect) introduceYourself() {
	fmt.Printf("Hi! My name is %v %v and I'm %v years old. My profession is %v %v. \n", a.name, a.lastName, a.age, a.kind, reflect.TypeOf(a).Name())
}

type profession interface {
	introduceYourself()
}

func introduceYourself(p profession) {
	p.introduceYourself()
	switch p.(type) {
	case dentist:
		fmt.Printf("I earn around %v per year.\n", p.(dentist).salary)
	case architect:
		fmt.Printf("My current projects are: %v\n", p.(architect).projects)
	}
}

func main() {

	d1 := dentist{
		person: person{
			name:     "Cristiano",
			lastName: "Ronaldo",
			age:      35,
		},
		salary:      50000,
		specialties: []string{"Pediatric dentistry", "Dental public health"},
	}

	d1.introduceYourself()

	a1 := architect{
		person: person{
			name:     "Lionel",
			lastName: "Messi",
			age:      33,
		},
		kind:     "Urban planning",
		projects: []string{"Project A", "Project B"},
	}

	a1.introduceYourself()

	fmt.Println("\n\t--> Using interface <--\n")

	introduceYourself(d1)
	introduceYourself(a1)

}
