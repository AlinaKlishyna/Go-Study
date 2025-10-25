package main

import "fmt"

// Обьявление переменных с указанием типа
var hi string
var a, b, c int
var max, min float64

// Обьявление и инициализация переменных без указания типа
var isActive = true
var str = "I'm good!"

// Для хранения символоclв принято использовать тип rune и int32
var symbol int32 = 'f'
var anotherSymbol rune = '?'

// Значения по умолчанию
var defaultString string // Пустая строка ""
var defaultInt int       // 0
var defaultFloat float64 // 0.0
var defaultBool bool     // false
var defaultRune rune     // 0

// Более краткий вывод переменной
var short int = 5
var anotherShort = 5

// Можно обновить сразу несколько переменных в блоке var
var (
	name string = "Alina" // Не нужно писать тип переменной (var)
	age  int    = 23
)

func main() {
	fmt.Println(str)
	fmt.Println(symbol)      // Выводит числовое значение символа
	fmt.Println(string(102)) // Преобразует числовое значение в символ

	fmt.Println(name, " - ", age)
}
