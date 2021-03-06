package main

import "fmt"

func main() {

	person := map[string]string{
		"Federer_Roger":     "Tennis",
		"Ronaldo_Cristiano": "Football",
		"Bolt_Usain":        "Sprinter",
	}

	for key, value := range person {
		fmt.Printf("Key: %v\tValue: %v\n", key, value)
	}

	fmt.Println("\n")
	person["Jordan_Michael"] = "Basketball"

	for key, value := range person {
		fmt.Printf("Key: %v\tValue: %v\n", key, value)
	}
	fmt.Println("\n")

	delete(person, "Federer_Roger")
	for key, value := range person {
		fmt.Printf("Key: %v\tValue: %v\n", key, value)
	}

}
