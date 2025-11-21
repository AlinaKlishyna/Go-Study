package database

import (
	"github.com/AlinaKlishyna/Go-Study.git/jwt/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const baseServer = "postgresql://postgres@localhost:5432/test"

func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(baseServer))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SetupDB() error {
	db, err := ConnectToDB()
	if err != nil {
		return err
	}

	if err := CreateTable(db, &core.User{}); err != nil {
		return err
	}

	return nil
}

func CreateTable(db *gorm.DB, dst ...interface{}) error {
	return db.AutoMigrate(dst...)
}
