package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("..\\Error-basic3\\names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
