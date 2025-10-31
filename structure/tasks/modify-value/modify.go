package main

import "fmt"

type User struct {
	FirstName string
}

// *User - место где в оперативной памяти эта структура хранится
// иначе изменит только копию
func setFirstName(firstName string, u *User) {
	u.FirstName = firstName
}

func main() {
	user := User{
		FirstName: "Alina",
	}

	setFirstName("Olga", &user) // не создается копия

	fmt.Printf("%+v", user)
}
