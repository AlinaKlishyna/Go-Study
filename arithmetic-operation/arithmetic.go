package main

import "fmt"

func main() {
	// := создаем переменные и присваиваем им значения
	// = присваиваем значение уже созданной переменной
	a := 5
	b := 10
	c := a + b
	fmt.Println("a = ", a, "\nb = ", b)
	fmt.Println("a + b =", c)

	c = a * b
	fmt.Println("a * b =", c)
	c = a - b
	fmt.Println("a - b =", c)
	c = a / b
	fmt.Println("a / b =", c)

	// При делении (/) если 2 целых числа, то результат будет округляться до целого
	var division int = 10 / 6 // 1
	fmt.Println("10 / 6 = ", division)

	// Чтобы получить дробное число при делении, как минимум один с операндов
	// должен быть дробным и сохранять результат так же должны в дробный тип
	var divisionFloat float32 = 10.0 / 6
	fmt.Println("10.0 / 6 = ", divisionFloat)

	// % - остаток от деления
	// участие могут принимать ТОЛЬКО целые числа

	// Формула: remainder = a - (a / b) * b, где (a / b) — целочисленное деление.
	// Пример: для 10 % 6: 10 / 6 = 1 => remainder = 10 - 1*6 = 4
	var remainder int = 10 % 6
	fmt.Println("10 % 6 = ", remainder)
	remainder = 10 % 3
	fmt.Println("10 % 3 = ", remainder)

	// Постфиксный инкремент и декремент
	var increment int = 0
	increment++ // increment = increment + 1
	fmt.Println("increment =", increment)
	increment--
	fmt.Println("decrement =", increment)
}
