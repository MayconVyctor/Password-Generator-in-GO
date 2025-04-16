package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var input int
	fmt.Println("Enter the number of characters of you password:")
	fmt.Scan(&input)
	password := generatePassword(input)
	fmt.Println(password)

}

func generatePassword(input int) string {

	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"

	allChars := lowerCase + upperCase + numbers + special

	// mentory := make([] byte, 2) {
	// 	lowerCase,
	// }

	password := make([]byte, input)

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}
	return string(password)
}
