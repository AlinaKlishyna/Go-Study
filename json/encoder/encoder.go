package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Car struct {
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Year     int    `json:"year"`
	Electric bool   `json:"electric,omitempty"`
}

func main() {
	car := Car{
		Brand:    "KIA",
		Model:    "Voll 200",
		Year:     2010,
		Electric: false,
	}

	Encoder(car, "cars.json")
	decoder := Decoder("cars.json")
	fmt.Println(decoder)
}

func Encoder(c Car, nameFile string) {
	// Создаём (или перезаписываем) файл
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Создаём энкодер, который будет писать прямо в файл
	encoder := json.NewEncoder(file)

	// Кодируем структуру и записываем в файл
	err = encoder.Encode(c)
	if err != nil {
		fmt.Println(err)
	}

	println("JSON записан в файл user.json")
}

func Decoder(nameFile string) Car {
	file, err := os.Open(nameFile)
	if err != nil {
		panic(err)
	}

	var car Car
	decoder := json.NewDecoder(file)
	decoder.Decode(&car)
	return car
}
