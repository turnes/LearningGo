package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(randomInt(1, 3))

	x := randomInt(0, 200)
	fmt.Printf("%v\n", x)

	x, s := randomIntString(0, 200, 10)
	fmt.Printf("%v\n", s)
}

func randomInt(min, max int) int {
	x := rand.Intn(max-min+1) + min
	return x
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(97, 122))
	}
	return string(bytes)
}

func randomIntString(min, max, len int) (int, string) {
	n := randomInt(min, max)
	s := randomString(len)
	return n, s

}
