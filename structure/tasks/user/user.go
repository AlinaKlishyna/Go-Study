package main

import "fmt"

type User struct {
	// i campi
	FirstName string
	LastName  string
	Birthday  int
}

// var user User, все поля структуры инициализируются значениями по умолчанию
// Для числовых типов это 0
// для строк — пустая строка
// для булевых — false
// для указателей и срезов — nil
func main() {
	user := User{"Alina", "Klishyna", 2001} // плохой вариант, тяжело расширять класс
	userAnother := User{
		FirstName: "Jonny",
		LastName:  "Fredrih",
		Birthday:  2014,
	}

	fmt.Println(user)               // {Alina Klishyna 2001}
	fmt.Printf("%v\n", userAnother) // {Jonny Fredrih 2014}
	fmt.Printf("%+v\n", user)       // {FirstName:Alina LastName:Klishyna Birthday:2001}

	var userZero User
	fmt.Printf("%+v\n", userZero) // {FirstName: LastName: Birthday:0}

	userZeroAnother := User{}
	fmt.Printf("%+v\n", userZeroAnother) // {FirstName: LastName: Birthday:0}

	// userOther := &User{}
	userOther := &User{}
	fmt.Printf("%+v\n", userOther) // &{FirstName: LastName: Birthday:0}
}
