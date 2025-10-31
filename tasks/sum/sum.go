package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	firstNumber := n % 10
	n /= 10
	secondNumber := n % 10
	n /= 10
	thirdNumber := n % 10

	sum := firstNumber + secondNumber + thirdNumber

	fmt.Println(sum)
}
