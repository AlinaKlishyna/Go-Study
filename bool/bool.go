package main

import "fmt"

func main() {
	var a int = 7
	var b int = 5
	// == Равно
	var c1 bool = a == b //
	var c2 bool = a == 7 // true

	// != Не равно
	var d1 bool = a != b // true
	var d2 bool = a != 8 // true

	// >= Больше или равно
	// <= Меньше или равно
	var e1 bool = a >= b // true
	var e2 bool = a <= b // false

	fmt.Println(c1, c2, d1, d2, e1, e2)

	if a < b {
		fmt.Println("a меньше b")
	} else {
		fmt.Println("a не меньше b")
	}

	// Логические операции
	// ! - НЕ (операция отрицания)
	d1 = !d2 // !true = false

	// Операция И - && (операция конъюнкции)
	// Возращает true если 2 true
	var result1 bool = (a > b) && (a == 7) // true && true = true
	var result2 bool = (a > b) && (a != 7) // true && false = false
	var result3 bool = (a < b) && (a != 7) // false && false = false

	// || - ИЛИ (операция дизъюнкции)
	// Возвращает true, если хотя бы 1 true
	var result4 bool = (a > b) || (a == 7) // true || true = true
	var result5 bool = (a > b) || (a != 7) // true || false = true
	var result6 bool = (a < b) || (a != 7) // false || false = false

	fmt.Println(d1)
	fmt.Println(result1, result2, result3)
	fmt.Println(result4, result5, result6)
}
