package main

import "fmt"

type client struct {
	name     string
	lastName string
	smoker   bool
}

func main() {

	client1 := client{
		name:     "LeBron",
		lastName: "James",
		smoker:   true,
	}

	fmt.Println(client1)

	client2 := client{"Lionel", "Messi", false}
	fmt.Println(client2)

}
