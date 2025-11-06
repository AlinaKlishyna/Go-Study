package main

import (
	"fmt"

	"gorm.io/driver/postgres" // go get -u gorm.io/driver/postgres - драйвер для Postgres
	"gorm.io/gorm"            // go get -u gorm.io/gorm - ORM библиотека
	// go get -u github.com/lib/pq - Postgres драйвер для database/sql
)

type User struct {
	gorm.Model // Встраиваемые поля (ID, CreatedAt, UpdatedAt, DeletedAt)
	Name       string
	LastName   string
	Age        int
}

func main() {
	db, err := gorm.Open(postgres.Open("postgresql://postgres@localhost:5432/postgres")) // Подключаемся к базе данных Postgres
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}) // Создаем таблицу User автоматически

	// Создаем новую запись
	err = db.Create(&User{ // Создаем новую запись в таблице User (строка INSERT)
		Name:     "Alina",
		LastName: "Klishyna",
		Age:      24,
	}).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS")

	var user User
	// Reading.. Where id = 3 AND last_name = 'Bidf'
	// db.First(&user, 13)                           // Ищем первую запись с primary key = 1
	db.First(&user, "last_name = ?", "Klishyna") // Ищем первую запись, где last_name = 'Bidf'
	fmt.Println("User - Where id = 3 AND last_name = 'Klishyna':", user)

	db.Model(&user).Update("last_name", "Krivorukov") // Обновляем поле last_name у найденной записи

	// меняет несколько, но не пустые
	db.Model(&user).Updates(User{ // Обновляем несколько полей у найденной записи
		Name:     "Metrd",
		LastName: "Petrov",
		Age:      30,
	})

	// База, вот точный список. Замени всё именно так, даже если я передаю пустую строку или ноль
	db.Model(&user).Updates(map[string]interface{}{ // Обновляем несколько полей у найденной записи через map
		"Name":     "Faf",
		"LastName": "Ivanova",
		"Age":      28,
	})

	fmt.Println("Updated User:", user)

	// db.Delete(&user) // Удаляем найденную запись
	// fmt.Println("Deleted User:", user)

	// Delete - delete user
	db.Delete(&user, "last_name = ?", "Klishyna")
}
