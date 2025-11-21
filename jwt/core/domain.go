package core

import (
	"github.com/AlinaKlishyna/Go-Study.git/jwt/core/resources"
	"gorm.io/gorm"
)

type Domain struct {
	db *gorm.DB
}

func NewDomain() (*Domain, error) {
	db, err := resources.ConnectToDB()
	if err != nil {
		panic(err)
	}

	return &Domain{db: db}, nil
}

func (d *Domain) DB() *gorm.DB {
	return d.db
}
