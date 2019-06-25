package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("BCrypt hashing example")
	password := "thisisnew378"
	hashedPassword, _ := generateHash(password)
	fmt.Println("hashed : ", hashedPassword)
	check := checkPassword(hashedPassword, "password")
	fmt.Println("Password match : ", check)
}

func generateHash(password string) (string, error) {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), error
}

func checkPassword(hashedPassword string, givenPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(givenPassword))
	return err == nil
}
