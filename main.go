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

	mandatory := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],
	}

	password := make([]byte, input-len(mandatory))

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	password = append(password, mandatory...)

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}
