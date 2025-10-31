package main

import "fmt"

type User struct {
	// i campi
	Name     Name
	Birthday string
	Age      int
}

type Name struct {
	// i campi
	FirstName string
	LastName  string
}

// func (к чему привязывается) название функции (что принимает) что возвращает {}
func (u User) Greet() {
	fmt.Printf("Sono %s %s, sono nata %s, ho %d anni. Piacere!\n",
		u.Name.FirstName,
		u.Name.LastName,
		u.Birthday,
		u.Age)
}

func main() {
	user := User{
		Name: Name{
			FirstName: "Alina",
			LastName:  "Klishyna",
		},
		Birthday: "24 gennaio",
		Age:      24,
	}

	fmt.Printf("%+v\n", user)

	User.Greet(user)
}
