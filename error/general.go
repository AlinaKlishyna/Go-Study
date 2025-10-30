package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Проверьте типы входных параметров.")
	} else {
		fmt.Println(divide(a, b))
	}

	// Создание собственной ошибки
	firstError := errors.New("My first error!")
	fmt.Println(firstError)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	}
	return a / b, nil
}
