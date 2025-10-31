package main

import (
	"errors"
	"fmt"
)

type Warrior struct {
	Type string
	Name string
}

func NewWarrior(name string) (*Warrior, error) {
	if name == "" {
		return nil, errors.New("empty name")

	}
	return &Warrior{
		Name: name,
		Type: "warrior",
	}, nil
}

func (w Warrior) Attack() string {
	return "Воин " + w.Name + " бьет мечом."
}

// -----------------------
type Mage struct {
	Type string
	Name string
}

func NewMage(name string) (*Mage, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Mage{
		Name: name,
		Type: "mage",
	}, nil
}

func (m Mage) Attack() string {
	return "Маг " + m.Name + " колдует огненный шар."
}

// -----------------------
type Archer struct {
	Type string
	Name string
}

func NewArcher(name string) (*Archer, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Archer{
		Name: name,
		Type: "archer",
	}, nil
}

func (a Archer) Attack() string {
	return "Лучник " + a.Name + " выпускает град стрел."
}

// --- Интерфейс персонажа ---
type Character interface {
	Attack() string
}

func Fight(c []Character) {
	for _, v := range c {
		fmt.Println(v.Attack())
	}
}

// --- Демонстрация работы ---
func main() {
	warrior, _ := NewWarrior("Король Артур")
	mage, _ := NewMage("Мерлин")
	archer, _ := NewArcher("Леголас")

	characters := []Character{warrior, mage, archer}
	Fight(characters)
}
