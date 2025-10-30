package main // Пакет main

import (
	"fmt"
	"structure/export-field/user" // Импортируем пакет user
)

func main() {
	user := user.User{
		FirstName: "Alina",
		LastName:  "Klishyna",
		Birthday:  25,
		// password:  "qwerty123", --- НЕВОЗМОЖНО ДОСТУПИТЬСЯ К ПОЛЮ password, ТАК КАК ОНО ЗАКРЫТОЕ вне ПАКЕТА user
	}

	fmt.Printf("%+v", user)
}
