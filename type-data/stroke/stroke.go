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
}
