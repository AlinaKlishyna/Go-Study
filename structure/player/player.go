package main

import "fmt"

// struttura
type Player struct {
	// i campi
	Name     Name
	Birthday int
	//struttura anonima
	Location struct {
		Adress  string
		Country string
	}
}

// структура именованная
type Name struct {
	FirstName string
	LastName  string
}

func main() {
	name := Name{
		FirstName: "Alina",
		LastName:  "Klishyna",
	}

	player := Player{
		Name:     name,
		Birthday: 2001,
		Location: struct { // Struttura anonima
			Adress  string
			Country string
		}{
			Adress:  "Via a Trato 14, Rovereto",
			Country: "Italia",
		},
	}

	fmt.Printf("%+v\n", player)        // {Name:{FirstName:Alina LastName:Klishyna} Birthday:2001}
	fmt.Println(player.Name.FirstName) // Alina

	player.Name = Name{
		FirstName: "Vlad",
		LastName:  "Lysiak",
	}

	player.Name.FirstName = "Bob"
	player.Name.LastName = "Stero"
}
