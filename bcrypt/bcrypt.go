package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for i := 0; i < 10; i++ {
		password := []byte("password")

		cost := 1
		hash, _ := bcrypt.GenerateFromPassword(password, cost)

		fmt.Printf("%s - %s\n", password, hash)
	}

	password := []byte("password")
	hashedPassword := "$2a$10$SxvAyTvuOXqpdiAMyIe1.OlRmIIn7ZjD2/FtocJeMzJo10VV9bHvu"
	check := check([]byte(hashedPassword), password)
	fmt.Println(check)

}

func check(hashedPassword []byte, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		return false
	}
	return true
}
