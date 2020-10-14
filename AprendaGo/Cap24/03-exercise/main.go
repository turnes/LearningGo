package main

import (
	"fmt"
)

type EspecialError struct {
	err string
}

func (e EspecialError) Error() string {
	return fmt.Sprintf("Printing from interface: %v", e.err)

}

func dealWithError(e error) {
	fmt.Println("Printing from function:\t", e)

}

func main() {

	myError := EspecialError{"I can't deal with it"}
	dealWithError(myError)

}
