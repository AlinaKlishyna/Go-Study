package main

import "fmt"

// Всякий раз, когда функция panic вступает в действие, она выполняет все функции defer, связанные с текущим потоком.
// Функции отложенного вызова defer могут применяться для освобождения ресурсов, использующихся в функции.
// Эти функции defer выполняются непосредственно перед завершением текущей функции.

// если есть panic и defer, то сначала выполняется defer, а потом panic
func main() {
	// Оператор defer позволяет выполнить определенную операцию после выполнения других действий
	// (даже если сработает panic), при этом не важно, где в реальности вызывается эта функция

	// будет выполнена в самом конце выполнения программы
	defer finish() // Отложенный вызов функции finish
	fmt.Println("Program has been started")
	fmt.Println("Program is working")

	// Program has been started
	// Program is working
	// Program has been finished

	// Порядок выполнения отложенных функций - в обратном порядке их объявления
	defer finish()                                // самый последний вызов
	defer fmt.Println("Program has been started") // передпоследний - Этот вызов отложен
	fmt.Println("Program is working")

	// Program is working
	// Program has been started
	// Program has been finished

	//Команда defer помещает вызов функции в стек. Поэтому вызовы функций выполняются в очередности
	// LIFO (Last-In, First-Out), то есть последняя отложенная функция будет выполнена первой

	// Оператор defer запоминает значения переменных, переданных в функцию, на момент объявления defer,
	// а не на момент её вызова. Это означает, что значение переменной будет захвачено сразу,
	// как только вы объявите отложенный вызов
	a := 5
	defer myFunc(a) // Когда вызовется myFunc, будет передано значение 5, а не 7
	a = 7
	fmt.Println("main - Value of a =", a)
	// main - Value of a = 7
	// myFunc - Value of a = 5
}

func finish() {
	fmt.Println("Program has been finished")

}

func myFunc(a int) {
	fmt.Println("myFunc - Value of a =", a)
}
