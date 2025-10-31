package main

import (
	"fmt"
)

type Event struct {
	Title    string
	Date     string
	Location string
}

func createGoEvent(title, date, location string) Event {
	event := Event{
		Title:    title,
		Date:     date,
		Location: location,
	}
	return event
}

func main() {
	event := createGoEvent("Alina", "24 january", "Rovereto")
	fmt.Println(event)
}
