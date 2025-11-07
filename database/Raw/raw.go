package main

import (
	"gorm.io/driver/postgres" // go get -u gorm.io/driver/postgres - драйвер для Postgres
	"gorm.io/gorm"
)

type Player struct {
	PlayerID uint   `gorm:"primaryKey"`
	Username string `gorm:"size:50"`
	Name     string `gorm:"size:50"`
	Age      int    `gprm:"size:3"`
	Email    string `gorm:"unique"`
}

func main() {
	// CONNECTION TO DATABASE
	db, err := gorm.Open(postgres.Open("postgresql://postgres@localhost:5432/postgres")) // Подключаемся к базе данных Postgres
	if err != nil {
		panic("failed to connect database")
	}

	db.Raw("SELECT id, name, age WHERE id = ?")
}
