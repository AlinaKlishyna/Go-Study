package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name       string
	LastName   string
	Age        int
	LicenseCar bool
}

func main() {
	user := User{
		Name:       "Alina",
		LastName:   "Klishyna",
		Age:        24,
		LicenseCar: false,
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// для чтения компьютером
	fmt.Printf("%s\n", data) // {"Name":"Alina","LastName":"Klishyna","Age":24,"LicenseCar":false}

	// MarshalIndent похож на Marshal, но применяет отступ (indent) для форматирования вывода
	dataIndent, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", dataIndent) // для чтения человеком
	//{
	// 	"Name": "Alina",
	//	"LastName": "Klishyna",
	// 	"Age": 24,
	//	"LicenseCar": false
	//}

	// Unmarshal - распаковка данных
	// переменные куда мы распаковываем данные
	var u User
	// json.Unmarshal(откуда([]byte)), куда (переменная структуры))
	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", u)

	// Проверка на является ли срез байтов форматом json  -   json.Valid()
	if !json.Valid(dataIndent) {
		fmt.Println("invalid json!") // вывод: invalid json!
	} else {
		fmt.Println("valid json!")
	}
}
