package main

import (
	"fmt"
	"reflect"
)

type User struct {
	// Теги — это метаданные, которые добавляются к полям структуры и могут использоваться для различных целей,
	// таких как сериализация, валидация и другие операции, требующие дополнительной информации о полях
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"username" db:"first_name"`
	Email     string `json:"email" db:"email"`
}

func main() {
	user := User{
		ID:        2,
		FirstName: "Alina",
		Email:     "book@gmail.com",
	}

	t := reflect.TypeOf(user)
	fmt.Println("Type: ", t.Name())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, JSON: %s, DB: %s\n",
			field.Name,
			field.Tag.Get("json"),
			field.Tag.Get("db"))
	}
}
