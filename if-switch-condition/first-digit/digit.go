package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)

	str := strconv.Itoa(a)
	firstChar := str[0]

	fmt.Println(string(firstChar))
}
