package main

import (
	"errors"
	"fmt"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrTooLowPasswordsCount = errors.New("too low passwords count")
	ErrTooBigPasswordsCount = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

func generatePassword(length int, count int) ([]string, error) {
	if length < MinPasswordLength {
		return nil, ErrPasswordLengthTooLow
	}
	if count < MinPasswordsCount {
		return nil, ErrTooLowPasswordsCount
	}
	if count > MaxPasswordsCount {
		return nil, ErrTooBigPasswordsCount
	}

	return nil, nil
}

func main() {
	passwords, err := generatePassword(12, 5)
	if err != nil {
		// Обработка ошибки

		// DA FARE
	}
	fmt.Println(passwords) // Вывод сгенерированных паролей
}
