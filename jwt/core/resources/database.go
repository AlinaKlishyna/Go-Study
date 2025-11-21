package resources

import (
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
