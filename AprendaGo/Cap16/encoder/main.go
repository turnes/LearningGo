package main

import (
	"encoding/json"
	"os"
)

type person struct {
	Name     string
	LastName string
	Age      int
	Address  string
	Hobbies  []string
}

func main() {

	p1 := person{
		Name:     "Cristiano",
		LastName: "Ronaldo",
		Age:      35,
		Address:  "The Stadium",
		Hobbies:  []string{"sport cars", "poker", "music"},
	}

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(p1)
}
