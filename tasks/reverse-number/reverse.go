package main

import "fmt"

// Дано трехзначное число. Переверните его, а затем выведите
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println("N number = ", n)

	firstNumber := (n % 10) * 100
	fmt.Println("First number = ", firstNumber)
	n /= 10
	fmt.Println("N number = ", n)

	secondNumber := (n % 10) * 10
	fmt.Println("Second number = ", secondNumber)
	n /= 10
	fmt.Println("N number = ", n)

	thirdNumber := n % 10
	fmt.Println("Third number = ", thirdNumber)

	n = firstNumber + secondNumber + thirdNumber
	fmt.Println(n)
}
