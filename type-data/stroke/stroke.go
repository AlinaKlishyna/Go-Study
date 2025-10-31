package main

import (
	"fmt"
	"strconv"
)

// Конвертация в строки
func main() {
	user := "ученик"
	steps := 4
	// fmt.Println("Поздравляю, " + user + "! Ты прошел " + steps + " шага по приведению типов") - Ошибка!! Разные типы
	fmt.Println("Поздравляю, " + user + "! Ты прошел " + strconv.Itoa(steps) + " шага по приведению типов")

	// strconv.Itoa()       int --> string      strconv.Itoa(5) --> "5"
	a := strconv.Itoa(2020) // int -> string
	fmt.Printf("%T \n", a)  // тип - string
	fmt.Println(a)          // 2020
	// !! Интересное дополнение, метод Itoa это всего-лишь обертка для FormatInt

	// strconv.Atoi()     string --> int      strconv.Atoi("10") --> 10
	strNum, _ := strconv.Atoi("555")         // string -> int    return int, err
	fmt.Printf("%d - %T \n", strNum, strNum) // 555 - int

	// Конвертация string в float с помощью метода ParseFloat -  ParseUint, ParseInt, ParseBool работают аналогично
	s := "23.23456"
	// как и в прошлом шаге, здесь 2 параметр - bitSize
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	// но нужно понимать что метод все равно вернет float64
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)         // 23.23456
	fmt.Printf("%T \n", result) // float64

	// Конкретный пример для разных bitSize:
	s = "1.0000000012345678"
	//  не будем обрабатывать ошибки в примерах, но на практике так не делайте ;)
	result32, _ := strconv.ParseFloat(s, 32)
	result64, _ := strconv.ParseFloat(s, 64)
	fmt.Println("bitSize32:", result32) // вывод 1 (не уместились)
	fmt.Println("bitSize64:", result64) // вывод  1.0000000012345678

	// Конвертация string в bool
	strtrue := "true"
	res, err := strconv.ParseBool(strtrue)
	if err != nil { // не забываем проверить ошибку
		panic(err)
	}
	fmt.Println(res)      // true
	fmt.Printf("%T", res) // bool
}
