package main

import (
	"errors"
	"fmt"
	"log"
)

// Композиция в Go позволяет создавать новые структуры
// включая в них другие структуры как поля
// Это позволяет создавать более сложные и гибкие структуры, используя уже существующие
type Car struct {
	Engine Engine
	Model  string
}

func (c *Car) Start() error {
	if err := c.Engine.Start(); err != nil {
		return fmt.Errorf("Engine start: %w", err)
	}
	// other services
	return nil
}

func (c *Car) Drive() error {
	if !c.Engine.Started {
		return errors.New("Engine not started!")
	}
	// ..
	return nil
}

type Engine struct {
	Started    bool
	HorsePower int
}

func (c *Engine) Start() error {
	if c.Started {
		return errors.New("Already started")
	}
	c.Started = true
	return nil
}

func main() {
	car := Car{
		Engine: Engine{
			Started:    false,
			HorsePower: 150,
		},
		Model: "KIA",
	}

	if err := car.Drive(); err != nil { // если ловим ошибку
		log.Fatal("Drive car: %v", err)
	}

	if err := car.Start(); err != nil { // если ловим ошибку
		log.Fatal("Start car: %v", err)
	}
	fmt.Printf("Car %s started..\n", car.Model)
	fmt.Printf("%+v\n", car)

	if err := car.Drive(); err != nil { // если ловим ошибку
		log.Fatal("Drive car: %v", err)
	}
}
