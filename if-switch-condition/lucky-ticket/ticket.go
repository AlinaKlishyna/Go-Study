package main

import "fmt"

func main() {
	var numberTicket int
	fmt.Scan(&numberTicket)

	var firstDigit = numberTicket % 10
	numberTicket /= 10
	var secondDigit = numberTicket % 10
	numberTicket /= 10
	var thirdDigit = numberTicket % 10
	numberTicket /= 10
	var fourthDigit = numberTicket % 10
	numberTicket /= 10
	var fifthDigit = numberTicket % 10
	numberTicket /= 10
	var sixthDigit = numberTicket % 10

	var sumFirst = firstDigit + secondDigit + thirdDigit
	var sumSecond = fourthDigit + fifthDigit + sixthDigit

	if sumFirst == sumSecond {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
