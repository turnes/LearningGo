package main

import (
	"encoding/json"
	"fmt"
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

	jsonOutput, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(p1)
	os.Stdout.Write(jsonOutput)
	//fmt.Println(string(jsonOutput))

}
