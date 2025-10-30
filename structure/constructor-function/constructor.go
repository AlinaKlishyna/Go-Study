package main

import (
	"fmt"
	"log"
	"structure/constructor-function/model"
)

func main() {
	model, err := model.NewPerson(
		"Alina",
		"Klishyna",
		24,
		"Via a Prato 14, Rovereto",
		"Italia",
	)

	if err != nil {
		log.Fatalf("Unable to create person: %v", err)
	}

	fmt.Printf("%+v", model)
}
