package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string   `json:"Name"`
	LastName string   `json:"LastName"`
	Age      int      `json:"Age"`
	Address  string   `json:"Address"`
	Hobbies  []string `json:"Free Time"`
}

func main() {
	sb := []byte(`{"Name":"Cristiano","LastName":"Ronaldo","Age":35,"Address":"The Stadium","Free Time":["sport cars","poker","music"]}`)

	var p1 person

	err := json.Unmarshal(sb, &p1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(p1)
}
