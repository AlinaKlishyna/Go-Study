package main

import "fmt"

// Обьявление переменных с указанием типа
var hi string
var a, b, c int
var max, min float64

// Обьявление и инициализация переменных без указания типа
var isActive = true
var str = "I'm good!"

// Для хранения символов принято использовать тип rune и int32
var symbol int32 = 'f'
var anotherSymbol rune = '?'

// Значения по умолчанию
var defaultString string // Пустая строка ""
var defaultInt int       // 0
var defaultFloat float64 // 0.0
var defaultBool bool     // false
var defaultRune rune     // 0

func main() {
	fmt.Println(str)
	fmt.Println(symbol)      // Выводит числовое значение символа
	fmt.Println(string(102)) // Преобразует числовое значение в символ
}
