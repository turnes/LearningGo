package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\t%v\n", x, string(x))
		}
	}

}
