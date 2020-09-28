package main

import (
	"fmt"
)

func say(word string) string {
	defer fmt.Println("defer in a function.") //3
	fmt.Println(word)                         //2
	return "returning something"

}

func main() {
	defer fmt.Println("Last thing done before a code block") //5
	fmt.Println("Hello defer")                               //1
	word := "Say something"
	word = say(word)
	fmt.Println(word) //4

}
