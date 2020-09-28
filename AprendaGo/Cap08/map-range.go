package main

import (
	"fmt"
)

func main() {
	protocol := map[int]string{
		21:  "ftp",
		22:  "ssh",
		25:  "smtp",
		53:  "dns",
		80:  "http",
		443: "https",
	}

	// not sequencial, execute severel times
	for number, name := range protocol {
		fmt.Printf("Port: %v\tProtocol: %v\n", number, name)
	}

	// delete
	fmt.Println("\n\tdeleting\n")
	delete(protocol, 25)
	for number, name := range protocol {
		fmt.Printf("Port: %v\tProtocol: %v\n", number, name)
	}

}
