package main

import "fmt"

func main() {

	menu := 3

	switch {
	case menu == 1:
		fmt.Println("Menu 1")
	case menu == 2:
		fmt.Println("Menu 2")
	case menu == 3:
		fmt.Println("Menu 3")
	case menu == 4:
		fmt.Println("Menu 4")
	}

}
