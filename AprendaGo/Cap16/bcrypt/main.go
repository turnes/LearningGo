package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "PassowordIsnotASafePassword"
	wrongpassword := "ThisIsNotTheRightPassword"
	cpassword, error := bcrypt.GenerateFromPassword([]byte(password), 10)

	if error != nil {
		fmt.Println("Error: ", error)
	}
	fmt.Println(string(cpassword))

	checkPass(cpassword, []byte(password))
	checkPass(cpassword, []byte(wrongpassword))

}

func checkPass(h, p []byte) {
	check := bcrypt.CompareHashAndPassword(h, p)
	if check == nil {
		fmt.Println("Password accepted")
	} else {
		fmt.Println("Wrong password")
	}

}
