package main

import "fmt"

const pi float64 = 3.1415

const (
	a float64 = 5
	b         // b = 5 - Принимает значение предыдущей константы
	c         // c = 5
)

// iota example
/* Этот пример аналогичен следующему коду
const(
	Monday = 0
	Tuesday = 1
	Wednesday = 2
	Thursday = 3
	Friday = 4
	Saturday = 5
	Sunday = 6
)
*/

// iota - используется в обьявлении констант для создания последовательных чисел
const (
	Monday    = iota // 0
	Tuesday          // 1
	Wednesday        // 2
	Thursday         // 3
	Friday           // 4
	Saturday         // 5
	Sunday           // 6
	_                // 7 - пропускаем значение
	add              // 8 - продолжение нумерации
)

// iota перечисляет по факту все константы в блоке
const (
	x float64 = iota * pi // 0 * pi = 0
	y         = iota * pi // 1 * pi = pi
	z         = iota * pi // 2 * pi = 2pi

	d = 5
	v
	j = iota // 5 - iota сбрасывается в новом блоке, а так она продолжает считать строки в блоке начиная с 0
)

// iota не меняется вне блока констант
const min = iota // 0
const max = iota // 0

func main() {
	fmt.Println(a, b, c) // 5 5 5

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday, add) // 0 1 2 3 4 5 6 8

	fmt.Println(x, y, z) // 0 3.1415 6.283
	fmt.Println(d, v, j) // 5 5 5
}
