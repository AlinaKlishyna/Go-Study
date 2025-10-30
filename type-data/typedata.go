package main

import (
	"fmt"
	"math"
	"strconv"
)

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
	fmt.Println(symbol) // Выводит числовое значение символа
	var letter = rune(102)
	fmt.Println(letter) // Преобразует числовое значение в символ

	fmt.Println(name, " - ", age)

	fmt.Println(str)
	fmt.Println(symbol) // Выводит числовое значение символа
	var letter = rune(102)
	fmt.Println(letter) // Преобразует числовое значение в символ

	fmt.Println(name, " - ", age)

	// Приведение целочисленных типов
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)

	fmt.Printf("%d - %T\n", bigIndex, bigIndex) // 15 - int32

	// Преобразование типов с меньшим количеством бит
	var big64 int64 = 64
	var little8 int8
	little8 = int8(big64)
	fmt.Println(little8) // 64

	// !! Однако важно помнить, что при преобразовании целых чисел может произойти потеря данных,
	// если результат превышает допустимый диапазон для целевого типа
	var big int64 = 129
	var little = int8(big)
	fmt.Println(little) // -127 (перенос данных)

	// Максимальные значения для целочисленных типов
	fmt.Println(math.MaxInt8)   // 127
	fmt.Println(math.MaxUint8)  // 255
	fmt.Println(math.MaxInt16)  // 32767
	fmt.Println(math.MaxUint16) // 65535

	fmt.Println(convert(342))

	// Приведение целых чисел и чисел с плавающей точкой
	var x int64 = 57
	var y float64 = float64(x)
	fmt.Println(y) // 57

	// Конвертация строк в байты и обратно
	str := "stroke"
	b := []byte(str) // "stroke" --> []byte
	c := string(b)   // [] byte --> "stroke"
	fmt.Println(str) // str
	fmt.Println(b)   // [115 116 114] - побайтовый срез
	fmt.Println(c)   // str

	// Конвертация строк в rune и обратно
	run := "строка"
	runb := []rune(run) // срез рун
	runc := string(runb)
	fmt.Println(run)  // строка
	fmt.Println(runb) // [1089 1090 1088 1086 1082 1072] - срез рун
	fmt.Println(runc) // строка

	// Конвертация bool в string
	var boolTest = true
	result := strconv.FormatBool(boolTest)
	fmt.Printf("%s - %T", result, result)
}

func convert(num int64) uint16 {
	return uint16(num)
}
