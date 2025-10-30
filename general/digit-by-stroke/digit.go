package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// fmt.Println(maxNumberByStroke("1234567890"))
	fmt.Println(squareDigitByStroke(9119))
}

func maxNumberByStroke(stroke string) int {
	array := []rune(stroke)
	max := int(array[0])
	for _, symbol := range array {
		if max < int(symbol) {
			max = int(symbol)
		}
	}

	// Go хранит символы не как цифры, а как их числовые коды (Unicode/ASCII)
	return max - '0'
}

func squareDigitByStroke(number int) int {
	var square string
	text := strconv.Itoa(number)
	for _, symbol := range text {
		digit := int(symbol - '0')
		square += strconv.Itoa(int(math.Pow(float64(digit), 2)))
	}
	result, _ := strconv.Atoi(square)
	return result
}
