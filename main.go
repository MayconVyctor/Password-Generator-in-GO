package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var input int
	fmt.Println("Enter the number of characters:")
	fmt.Scan(&input)
	password := generatePassword(input)
	fmt.Println(password)

}

func generatePassword(input int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()+?><:{}[]"

	// lowerCase := "abcdefghijklmnopqrstuvwxyz"
	// upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// numbers := "0123456789"
	// special := "!@#$%^&*()+?><:{}[]"

	// allchars := lowerCase, upperCase, numbers, special

	password := make([]byte, input)

	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}
