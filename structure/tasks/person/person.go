package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Bla-bla-bla, my name is " + p.Name)
}

type Andriod struct {
	Person
	Model string
}

func main() {
	a := Andriod{
		Model: "IY-600",
		Person: Person{
			Name: "Alina",
		},
	}
	fmt.Println(a)
}
