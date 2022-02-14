package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	stringConnection := "host=localhost port=5432 user=root dbname=root password=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("failed to connect database")
	}
}
